package valid8_test

import (
	"testing"

	"codeberg.org/audryus/valid8"
)

func TestFRTranslations(t *testing.T) {
	v := valid8.New(valid8.WithLocales(valid8.FR))
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

		// Other tags
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
		PostcodeIso3166Alpha2  string `validate:"postcode_iso3166_alpha2=FR"`
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
		CountryCode:  "FR",

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
		E164:                   "abc",
		UrnRFC2141:             "abc",
		Fqdn:                   "abc",
		Unique:                 []int{1, 1},
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

	err := v.Struct(s, valid8.FR)
	if len(err) == 0 {
		t.Fatal("Expected validation errors, got none")
	}

	errs := valid8.ErrorsToMap(err)

	// Required variants
	containsEquals(t, errs, "teststruct.requiredif", "RequiredIf est un champ requis")
	containsEquals(t, errs, "teststruct.requiredunless", "RequiredUnless est un champ requis")
	containsEquals(t, errs, "teststruct.requiredwith", "RequiredWith est un champ requis")
	containsEquals(t, errs, "teststruct.requiredwithall", "RequiredWithAll est un champ requis")
	containsEquals(t, errs, "teststruct.requiredwithout", "RequiredWithout est un champ requis")
	containsEquals(t, errs, "teststruct.requiredwithoutall", "RequiredWithoutAll est un champ requis")

	// Excluded variants
	containsEquals(t, errs, "teststruct.excludedif", "ExcludedIf est un champ exclu")
	containsEquals(t, errs, "teststruct.excludedunless", "ExcludedUnless est un champ exclu")
	containsEquals(t, errs, "teststruct.excludedwith", "ExcludedWith est un champ exclu")
	containsEquals(t, errs, "teststruct.excludedwithall", "ExcludedWithAll est un champ exclu")
	containsEquals(t, errs, "teststruct.excludedwithout", "ExcludedWithout est un champ exclu")
	containsEquals(t, errs, "teststruct.excludedwithoutall", "ExcludedWithoutAll est un champ exclu")

	// Other tags
	containsEquals(t, errs, "teststruct.isdefault", "IsDefault doit être dans la valeur par défaut")
	containsEquals(t, errs, "teststruct.alphaspace", "AlphaSpace doit contenir uniquement des caractères alphabétiques et des espaces")
	containsEquals(t, errs, "teststruct.alphanumspace", "AlphanumSpace doit contenir uniquement des caractères alphanumériques et des espaces")
	containsEquals(t, errs, "teststruct.alphaunicode", "AlphaUnicode doit contenir uniquement des caractères unicode alphabétiques")
	containsEquals(t, errs, "teststruct.alphanumunicode", "AlphanumUnicode doit contenir uniquement des caractères unicode alphanumériques")
	containsEquals(t, errs, "teststruct.e164", "E164 doit être un numéro de téléphone au format E.164 valide")
	containsEquals(t, errs, "teststruct.urnrfc2141", "UrnRFC2141 doit être un URN RFC 2141 valide")
	containsEquals(t, errs, "teststruct.fqdn", "Fqdn doit être un FQDN valide")
	containsEquals(t, errs, "teststruct.unique", "Unique doit contenir des valeurs uniques")
	containsEquals(t, errs, "teststruct.cron", "Cron doit être une expression cron valide")
	containsEquals(t, errs, "teststruct.json", "Json doit être une chaîne json valide")
	containsEquals(t, errs, "teststruct.jwt", "Jwt doit être une chaîne jwt valide")
	containsEquals(t, errs, "teststruct.lowercase", "Lowercase doit être une chaîne en minuscules")
	containsEquals(t, errs, "teststruct.uppercase", "Uppercase doit être une chaîne en majuscules")
	containsEquals(t, errs, "teststruct.datetime", "Datetime ne correspond pas au format 2006-01-02")
	containsEquals(t, errs, "teststruct.timezone", "Timezone doit être un fuseau horaire valide")
	containsEquals(t, errs, "teststruct.postcodeiso3166alpha2", "PostcodeIso3166Alpha2 ne correspond pas au format de code postal du pays FR")
	containsEquals(t, errs, "teststruct.postcodeiso3166alpha2f", "PostcodeIso3166Alpha2F ne correspond pas au format de code postal du champ CountryCode")
	containsEquals(t, errs, "teststruct.bcp47languagetag", "Bcp47LanguageTag doit être un code de langue BCP 47 valide")
}
