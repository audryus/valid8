package valid8_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/audryus/valid8"
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
