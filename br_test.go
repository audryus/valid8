package valid8_test

import (
	"testing"

	"codeberg.org/audryus/valid8"
)

func TestPTBRTranslations(t *testing.T) {
	v := valid8.New(valid8.WithLocales(valid8.PT_BR))
	if v == nil {
		t.Fatal("New() returned nil")
	}

	type TestStruct struct {
		// Required variants
		RequiredIf         string `validate:"required_if=OtherField value"`
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

		// Other custom tags
		IsDefault              int    `validate:"isdefault"`
		AlphaSpace             string `validate:"alphaspace"`
		AlphanumSpace          string `validate:"alphanumspace"`
		AlphaUnicode           string `validate:"alphaunicode"`
		AlphanumUnicode        string `validate:"alphanumunicode"`
		E164                   string `validate:"e164"`
		UrnRFC2141             string `validate:"urn_rfc2141"`
		Fqdn                   string `validate:"fqdn"`
		Unique                 []int  `validate:"unique"`
		Cron                   string `validate:"cron"`
		Json                   string `validate:"json"`
		Jwt                    string `validate:"jwt"`
		Lowercase              string `validate:"lowercase"`
		Uppercase              string `validate:"uppercase"`
		Datetime               string `validate:"datetime=2006-01-02"`
		Timezone               string `validate:"timezone"`
		PostcodeIso3166Alpha2  string `validate:"postcode_iso3166_alpha2=BR"`
		PostcodeIso3166Alpha2F string `validate:"postcode_iso3166_alpha2_field=CountryCode"`
		Bcp47LanguageTag       string `validate:"bcp47_language_tag"`

		// Helper fields
		OtherField          string
		AnotherField        string
		CountryCode         string
		MissingField        string
		AnotherMissingField string
	}

	s := TestStruct{
		// Trigger required_... (mostly by leaving them empty and setting helpers)
		OtherField:   "value",
		AnotherField: "value",
		CountryCode:  "BR",
		// MissingField stays empty

		// Trigger excluded_... (by setting them while helpers are present)
		ExcludedIf:         "present",
		ExcludedUnless:     "present", // OtherField is "value", which is not "other"
		ExcludedWith:       "present",
		ExcludedWithAll:    "present",
		ExcludedWithout:    "present", // MissingField is empty
		ExcludedWithoutAll: "present", // Both missing are empty

		// Trigger others with invalid data
		IsDefault:              1,
		AlphaSpace:             "123",
		AlphanumSpace:          "!",
		AlphaUnicode:           "123",
		AlphanumUnicode:        "!",
		E164:                   "abc",
		UrnRFC2141:             "abc",
		Fqdn:                   "abc",
		Unique:                 []int{1, 1},
		Cron:                   "abc",
		Json:                   "abc",
		Jwt:                    "abc",
		Lowercase:              "ABC",
		Uppercase:              "abc",
		Datetime:               "2023/01/01", // Wrong format
		Timezone:               "abc",
		PostcodeIso3166Alpha2:  "abc",
		PostcodeIso3166Alpha2F: "abc",
		Bcp47LanguageTag:       "!!!",
	}

	err := v.Struct(s, valid8.PT_BR)
	if len(err) == 0 {
		t.Fatal("Expected validation errors, got none")
	}

	errs := valid8.ErrorsToMap(err)

	// Required variants
	containsEquals(t, errs, "teststruct.requiredif", "RequiredIf é um campo obrigatório")
	containsEquals(t, errs, "teststruct.requiredunless", "RequiredUnless é um campo obrigatório")
	containsEquals(t, errs, "teststruct.requiredwith", "RequiredWith é um campo obrigatório")
	containsEquals(t, errs, "teststruct.requiredwithall", "RequiredWithAll é um campo obrigatório")
	containsEquals(t, errs, "teststruct.requiredwithout", "RequiredWithout é um campo obrigatório")
	containsEquals(t, errs, "teststruct.requiredwithoutall", "RequiredWithoutAll é um campo obrigatório")

	// Excluded variants
	containsEquals(t, errs, "teststruct.excludedif", "ExcludedIf é um campo excluído")
	containsEquals(t, errs, "teststruct.excludedunless", "ExcludedUnless é um campo excluído")
	containsEquals(t, errs, "teststruct.excludedwith", "ExcludedWith é um campo excluído")
	containsEquals(t, errs, "teststruct.excludedwithall", "ExcludedWithAll é um campo excluído")
	containsEquals(t, errs, "teststruct.excludedwithout", "ExcludedWithout é um campo excluído")
	containsEquals(t, errs, "teststruct.excludedwithoutall", "ExcludedWithoutAll é um campo excluído")

	// Other tags
	containsEquals(t, errs, "teststruct.isdefault", "IsDefault deve estar no valor padrão")
	containsEquals(t, errs, "teststruct.alphaspace", "AlphaSpace deve conter apenas caracteres alfabéticos e espaços")
	containsEquals(t, errs, "teststruct.alphanumspace", "AlphanumSpace deve conter apenas caracteres alfanuméricos e espaços")
	containsEquals(t, errs, "teststruct.alphaunicode", "AlphaUnicode deve conter apenas caracteres unicode alfabéticos")
	containsEquals(t, errs, "teststruct.alphanumunicode", "AlphanumUnicode deve conter apenas caracteres unicode alfanuméricos")
	containsEquals(t, errs, "teststruct.e164", "E164 deve ser um número de telefone no formato E.164 válido")
	containsEquals(t, errs, "teststruct.urnrfc2141", "UrnRFC2141 deve ser um URN RFC 2141 válido")
	containsEquals(t, errs, "teststruct.fqdn", "Fqdn deve ser um FQDN válido")
	containsEquals(t, errs, "teststruct.unique", "Unique deve conter valores únicos")
	containsEquals(t, errs, "teststruct.cron", "Cron deve ser uma expressão cron válida")
	containsEquals(t, errs, "teststruct.json", "Json deve ser uma string json válida")
	containsEquals(t, errs, "teststruct.jwt", "Jwt deve ser uma string jwt válida")
	containsEquals(t, errs, "teststruct.lowercase", "Lowercase deve ser uma string em minúsculo")
	containsEquals(t, errs, "teststruct.uppercase", "Uppercase deve ser uma string em maiúsculo")
	containsEquals(t, errs, "teststruct.datetime", "Datetime não corresponde ao formato 2006-01-02")
	containsEquals(t, errs, "teststruct.timezone", "Timezone deve ser um fuso horário válido")
	containsEquals(t, errs, "teststruct.postcodeiso3166alpha2", "PostcodeIso3166Alpha2 não corresponde ao formato de CEP do país BR")
	containsEquals(t, errs, "teststruct.postcodeiso3166alpha2f", "PostcodeIso3166Alpha2F não corresponde ao formato de CEP do campo CountryCode")
	containsEquals(t, errs, "teststruct.bcp47languagetag", "Bcp47LanguageTag deve ser um código de idioma BCP 47 válido")
}
