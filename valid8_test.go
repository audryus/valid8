package valid8_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/audryus/valid8"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func containsEquals(t *testing.T, mep map[string]string, key, value string) {
	v, ok := mep[key]
	if !ok {
		t.Errorf("Key %s not found", key)
	} else if v != value {
		t.Errorf("Founded %s, expected %s", v, value)
	}
}

func TestRegisterTranslation(t *testing.T) {
	v := valid8.New(valid8.WithLocales(valid8.PT_BR))

	// Register a new validation tag that always fails
	err := v.Validator.RegisterValidation("mytag", func(fl validator.FieldLevel) bool {
		return false
	})
	if err != nil {
		t.Fatalf("Failed to register validation: %v", err)
	}

	// Register the translation for PT_BR without explicitly providing TranslationFn.
	// It should now fallback to DefaultTranslationFn automatically.
	v.RegisterTranslation(valid8.Translation{
		Tag:    "mytag",
		Text:   "{0} invalida",
		Locale: valid8.PT_BR,
	})

	type TestStruct struct {
		Field string `validate:"mytag"`
	}

	s := TestStruct{Field: "anything"}

	// Test PT_BR
	errs := v.Struct(s, valid8.PT_BR)
	mep := valid8.ErrorsToMap(errs)
	containsEquals(t, mep, "teststruct.field", "Field invalida")

	// Test ES (uses fallback EN)
	errs = v.Struct(s, valid8.ES)
	mep = valid8.ErrorsToMap(errs)
	containsEquals(t, mep, "teststruct.field", "Field invalida")

	// Register the translation for ES without explicitly providing TranslationFn.
	// It should now fallback to DefaultTranslationFn automatically.
	v.RegisterTranslation(valid8.Translation{
		Tag:    "mytag",
		Text:   "{0} no es valido",
		Locale: valid8.ES,
	})

	// Test ES
	errs = v.Struct(s, valid8.ES)
	mep = valid8.ErrorsToMap(errs)
	containsEquals(t, mep, "teststruct.field", "Field no es valido")

	// Test DE (uses fallback EN)
	errs = v.Struct(s, valid8.DE)
	mep = valid8.ErrorsToMap(errs)
	containsEquals(t, mep, "teststruct.field", "Field no es valido")

	// Register the translation for DE without explicitly providing TranslationFn.
	// It should now fallback to DefaultTranslationFn automatically.
	v.RegisterTranslation(valid8.Translation{
		IgnoreFallback: true,
		Tag:            "mytag",
		Text:           "{0} ist nicht gültig",
		Locale:         valid8.DE,
	})

	// Test FR (uses fallback EN)
	errs = v.Struct(s, valid8.FR)
	mep = valid8.ErrorsToMap(errs)
	containsEquals(t, mep, "teststruct.field", "Field no es valido")
}

func TestRegisterTranslationErrors(t *testing.T) {
	v := valid8.New()

	tests := []struct {
		name        string
		translation valid8.Translation
		expectedErr error
	}{
		{
			name: "Missing Tag",
			translation: valid8.Translation{
				Text:   "some text",
				Locale: valid8.EN,
			},
			expectedErr: valid8.ErrTagMandatory,
		},
		{
			name: "Missing Text",
			translation: valid8.Translation{
				Tag:    "required",
				Locale: valid8.EN,
			},
			expectedErr: valid8.ErrTextMandatory,
		},
		{
			name: "Missing Locale",
			translation: valid8.Translation{
				Tag:  "required",
				Text: "is required",
			},
			expectedErr: valid8.ErrLocaleMandatory,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := v.RegisterTranslation(tt.translation)
			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("Expected error %v, got %v", tt.expectedErr, err)
			}
		})
	}
}

func TestWithAllLocales(t *testing.T) {
	v := valid8.New(valid8.WithAllLocales())
	if v == nil {
		t.Fatal("New() returned nil")
	}

	type TestStruct struct {
		Name string `validate:"required"`
	}

	s := TestStruct{Name: ""}

	tests := []struct {
		lang     valid8.Locale
		expected string
	}{
		{valid8.EN, "Name is a required field"},
		{valid8.PT_BR, "Name é um campo obrigatório"},
		{valid8.ES, "Name es un campo requerido"},
		{valid8.FR, "Name est un champ obligatoire"},
		{valid8.DE, "Name ist ein Pflichtfeld"},
		{valid8.JA, "Nameは必須フィールドです"},
		{valid8.ZH, "Name为必填字段"},
		{valid8.ZHTW, "Name為必填欄位"},
	}

	for _, tt := range tests {
		t.Run(string(tt.lang), func(t *testing.T) {
			errs := v.Struct(s, tt.lang)
			if len(errs) == 0 {
				t.Errorf("Expected errors for locale %s, got none", tt.lang)
				return
			}
			mep := valid8.ErrorsToMap(errs)
			containsEquals(t, mep, "teststruct.name", tt.expected)
		})
	}
}

func TestRegisterTranslationAdvanced(t *testing.T) {
	v := valid8.New(valid8.WithLocales(valid8.PT_BR))

	// Test Custom TranslationFn
	v.RegisterTranslation(valid8.Translation{
		Tag:    "customfn",
		Text:   "ignored text",
		Locale: valid8.PT_BR,
		TranslationFn: func(ut ut.Translator, fe validator.FieldError) string {
			return "COMPLETELY CUSTOM"
		},
	})

	type TestStruct struct {
		Field string `validate:"customfn"`
	}
	v.Validator.RegisterValidation("customfn", func(fl validator.FieldLevel) bool { return false })

	errs := v.Struct(TestStruct{Field: "x"}, valid8.PT_BR)
	if len(errs) != 1 || errs[0].Message() != "COMPLETELY CUSTOM" {
		t.Errorf("Expected COMPLETELY CUSTOM, got %v", errs[0].Message())
	}

	// Test IgnoreFallback when locale is NOT registered
	// FR is not registered, IgnoreFallback is true.
	// This should result in no translation being registered at all (since EN fallback is skipped).
	v.RegisterTranslation(valid8.Translation{
		Tag:            "onlyfr",
		Text:           "seulement fr",
		Locale:         valid8.FR,
		IgnoreFallback: true,
	})
	v.Validator.RegisterValidation("onlyfr", func(fl validator.FieldLevel) bool { return false })

	errs = v.Struct(struct {
		Field string `validate:"onlyfr"`
	}{Field: "x"}, valid8.EN)
	// When translation is missing, it returns the tag or a default message from validator
	if len(errs) != 1 || errs[0].Message() == "seulement fr" {
		t.Errorf("Expected translation to be missing, but got %v", errs[0].Message())
	}
}

func TestStructAdvanced(t *testing.T) {
	v := valid8.New()

	// Test no errors
	type TestStruct struct {
		Field string `validate:"required"`
	}
	errs := v.Struct(TestStruct{Field: "valid"}, valid8.EN)
	if len(errs) != 0 {
		t.Errorf("Expected no errors, got %d", len(errs))
	}

	// Test locale not found, fallback to EN
	// PT_BR is not registered, so it should use EN translation for 'required'
	errs = v.Struct(TestStruct{Field: ""}, valid8.PT_BR)
	if len(errs) == 0 {
		t.Fatal("Expected errors, got none")
	}
	mep := valid8.ErrorsToMap(errs)
	containsEquals(t, mep, "teststruct.field", "Field is a required field")
}

func TestRegisterTranslationOverride(t *testing.T) {
	v := valid8.New()

	v.Validator.RegisterValidation("ovr", func(fl validator.FieldLevel) bool { return false })

	// First registration
	v.RegisterTranslation(valid8.Translation{
		Tag:    "ovr",
		Text:   "{0} Initial",
		Locale: valid8.EN,
	})

	// Second registration (overrides)
	v.RegisterTranslation(valid8.Translation{
		Tag:    "ovr",
		Text:   "{0} Overridden",
		Locale: valid8.EN,
	})

	type TestStruct struct {
		Field string `validate:"ovr"`
	}
	errs := v.Struct(TestStruct{Field: "x"}, valid8.EN)
	if errs[0].Message() != "Field Overridden" {
		t.Errorf("Expected Field Overridden, got %v", errs[0].Message())
	}

	// Third registration (IgnoreOverride = true)
	v.RegisterTranslation(valid8.Translation{
		Tag:            "ovr",
		Text:           "{0} Ignored",
		Locale:         valid8.EN,
		IgnoreOverride: true,
	})

	errs = v.Struct(TestStruct{Field: "x"}, valid8.EN)
	if errs[0].Message() != "Field Overridden" {
		t.Errorf("Expected Field Overridden (no change), got %v", errs[0].Message())
	}
}

func TestHTTPIntegration(t *testing.T) {
	v := valid8.New(valid8.WithAllLocales())

	// Mock server handler
	handler := func(w http.ResponseWriter, r *http.Request) {
		lang := valid8.Locale(r.Header.Get("Accept-Language"))

		type Request struct {
			Email string `validate:"required,email"`
		}

		req := Request{Email: "invalid-email"} // Trigger email error
		errs := v.Struct(req, lang)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(valid8.ErrorsToMap(errs))
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	tests := []struct {
		lang     string
		expected string
	}{
		{"pt_BR", "Email deve ser um endereço de e-mail válido"},
		{"en", "Email must be a valid email address"},
		{"es", "Email debe ser una dirección de correo electrónico válida"},
	}

	for _, tt := range tests {
		t.Run(tt.lang, func(t *testing.T) {
			req, _ := http.NewRequest("GET", server.URL, nil)
			req.Header.Set("Accept-Language", tt.lang)

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Fatal(err)
			}
			defer resp.Body.Close()

			var result map[string]string
			json.NewDecoder(resp.Body).Decode(&result)

			containsEquals(t, result, "request.email", tt.expected)
		})
	}
}
