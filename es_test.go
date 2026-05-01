package valid8_test

import (
	"testing"

	"github.com/audryus/valid8"
)

func TestESTranslations(t *testing.T) {
	v := valid8.New(valid8.WithLocales(valid8.ES))
	if v == nil {
		t.Fatal("New() returned nil")
	}

	type TestStruct struct {
		// Required variants
		RequiredUnless     string `validate:"required_unless=OtherField other"`
		RequiredWith       string `validate:"required_with=OtherField"`
		RequiredWithAll    string `validate:"required_with_all=OtherField AnotherField"`
		RequiredWithout    string `validate:"required_without=MissingField"`
		RequiredWithoutAll string `validate:"required_without_all=MissingField AnotherMissingField"`

		// Excluded variants
		ExcludedIf         string `validate:"excluded_if=OtherField value"`
		ExcludedUnless     string `validate:"excluded_unless=OtherField other"`
		ExcludedWith       string `validate:"excluded_with=OtherField"`
		ExcludedWithAll    string `validate:"excluded_with_all=OtherField AnotherField"`
		ExcludedWithout    string `validate:"excluded_without=MissingField"`
		ExcludedWithoutAll string `validate:"excluded_without_all=MissingField AnotherMissingField"`

		// Other tags
		IsDefault               int      `validate:"isdefault"`
		AlphaSpace              string   `validate:"alphaspace"`
		AlphanumSpace           string   `validate:"alphanumspace"`
		AlphaUnicode            string   `validate:"alphaunicode"`
		AlphanumUnicode         string   `validate:"alphanumunicode"`
		UrnRFC2141              string   `validate:"urn_rfc2141"`
		Fqdn                    string   `validate:"fqdn"`
		Cron                    string   `validate:"cron"`
		Json                    string   `validate:"json"`
		Jwt                     string   `validate:"jwt"`
		Lowercase               string   `validate:"lowercase"`
		Uppercase               string   `validate:"uppercase"`
		Datetime                string   `validate:"datetime=2006-01-02"`
		Timezone                string   `validate:"timezone"`
		PostcodeIso3166Alpha2   string   `validate:"postcode_iso3166_alpha2=ES"`
		PostcodeIso3166Alpha2F  string   `validate:"postcode_iso3166_alpha2_field=CountryCode"`
		Bcp47LanguageTag        string   `validate:"bcp47_language_tag"`

		// Helper fields
		OtherField           string
		AnotherField         string
		CountryCode          string
		MissingField         string
		AnotherMissingField  string
	}

	s := TestStruct{
		// Trigger required_... (mostly by leaving them empty and setting helpers)
		OtherField:   "value",
		AnotherField: "value",
		CountryCode:  "ES",

		// Trigger excluded_... (by setting them while helpers are present)
		ExcludedIf:         "present",
		ExcludedUnless:     "present",
		ExcludedWith:       "present",
		ExcludedWithAll:    "present",
		ExcludedWithout:    "present",
		ExcludedWithoutAll: "present",

		// Trigger others with invalid data
		IsDefault:              1,
		AlphaSpace:             "123",
		AlphanumSpace:          "!",
		AlphaUnicode:           "123",
		AlphanumUnicode:        "!",
		UrnRFC2141:             "abc",
		Fqdn:                   "abc",
		Cron:                   "abc",
		Json:                   "abc",
		Jwt:                    "abc",
		Lowercase:              "ABC",
		Uppercase:              "abc",
		Datetime:               "2023/01/01",
		Timezone:               "abc",
		PostcodeIso3166Alpha2:  "abc",
		PostcodeIso3166Alpha2F: "abc",
		Bcp47LanguageTag:       "!!!",
	}

	err := v.Struct(s, valid8.ES)
	if len(err) == 0 {
		t.Fatal("Expected validation errors, got none")
	}

	errs := valid8.ErrorsToMap(err)

	// Required variants
	containsEquals(t, errs, "teststruct.requiredunless", "RequiredUnless es un campo obligatorio")
	containsEquals(t, errs, "teststruct.requiredwith", "RequiredWith es un campo obligatorio")
	containsEquals(t, errs, "teststruct.requiredwithall", "RequiredWithAll es un campo obligatorio")
	containsEquals(t, errs, "teststruct.requiredwithout", "RequiredWithout es un campo obligatorio")
	containsEquals(t, errs, "teststruct.requiredwithoutall", "RequiredWithoutAll es un campo obligatorio")

	// Excluded variants
	containsEquals(t, errs, "teststruct.excludedif", "ExcludedIf es un campo excluido")
	containsEquals(t, errs, "teststruct.excludedunless", "ExcludedUnless es un campo excluido")
	containsEquals(t, errs, "teststruct.excludedwith", "ExcludedWith es un campo excluido")
	containsEquals(t, errs, "teststruct.excludedwithall", "ExcludedWithAll es un campo excluido")
	containsEquals(t, errs, "teststruct.excludedwithout", "ExcludedWithout es un campo excluido")
	containsEquals(t, errs, "teststruct.excludedwithoutall", "ExcludedWithoutAll es un campo excluido")

	// Other tags
	containsEquals(t, errs, "teststruct.isdefault", "IsDefault debe estar en el valor predeterminado")
	containsEquals(t, errs, "teststruct.alphaspace", "AlphaSpace debe contener solo caracteres alfabéticos y espacios")
	containsEquals(t, errs, "teststruct.alphanumspace", "AlphanumSpace debe contener solo caracteres alfanuméricos y espacios")
	containsEquals(t, errs, "teststruct.alphaunicode", "AlphaUnicode debe contener solo caracteres unicode alfabéticos")
	containsEquals(t, errs, "teststruct.alphanumunicode", "AlphanumUnicode debe contener solo caracteres unicode alfanuméricos")
	containsEquals(t, errs, "teststruct.urnrfc2141", "UrnRFC2141 debe ser un URN RFC 2141 válido")
	containsEquals(t, errs, "teststruct.fqdn", "Fqdn debe ser un FQDN válido")
	containsEquals(t, errs, "teststruct.cron", "Cron debe ser una expresión cron válida")
	containsEquals(t, errs, "teststruct.json", "Json debe ser una cadena json válida")
	containsEquals(t, errs, "teststruct.jwt", "Jwt debe ser una cadena jwt válida")
	containsEquals(t, errs, "teststruct.lowercase", "Lowercase debe ser una cadena en minúsculas")
	containsEquals(t, errs, "teststruct.uppercase", "Uppercase debe ser una cadena en mayúsculas")
	containsEquals(t, errs, "teststruct.datetime", "Datetime no coincide con el formato 2006-01-02")
	containsEquals(t, errs, "teststruct.timezone", "Timezone debe ser una zona horaria válida")
	containsEquals(t, errs, "teststruct.postcodeiso3166alpha2", "PostcodeIso3166Alpha2 no coincide con el formato de código postal del país ES")
	containsEquals(t, errs, "teststruct.postcodeiso3166alpha2f", "PostcodeIso3166Alpha2F no coincide con el formato de código postal del campo CountryCode")
	containsEquals(t, errs, "teststruct.bcp47languagetag", "Bcp47LanguageTag debe ser un código de idioma BCP 47 válido")
}
