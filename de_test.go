package valid8_test

import (
	"testing"

	"codeberg.org/audryus/valid8"
)

func TestDETranslations(t *testing.T) {
	v := valid8.New(valid8.WithLocales(valid8.DE))
	if v == nil {
		t.Fatal("New() returned nil")
	}

	type TestStruct struct {
		AlphaSpace       string `validate:"alphaspace"`
		AlphanumSpace    string `validate:"alphanumspace"`
		AlphaUnicode     string `validate:"alphaunicode"`
		AlphanumUnicode  string `validate:"alphanumunicode"`
		UrnRFC2141       string `validate:"urn_rfc2141"`
		Timezone         string `validate:"timezone"`
		Bcp47LanguageTag string `validate:"bcp47_language_tag"`
		ValidateFn       string `validate:"validateFn"`
	}

	s := TestStruct{
		AlphaSpace:       "123",
		AlphanumSpace:    "!",
		AlphaUnicode:     "123",
		AlphanumUnicode:  "!",
		UrnRFC2141:       "abc",
		Timezone:         "abc",
		Bcp47LanguageTag: "!!!",
		ValidateFn:       "invalid",
	}

	err := v.Struct(s, valid8.DE)
	if len(err) == 0 {
		t.Fatal("Expected validation errors, got none")
	}

	errs := valid8.ErrorsToMap(err)

	containsEquals(t, errs, "teststruct.alphaspace", "AlphaSpace darf nur Buchstaben und Leerzeichen enthalten")
	containsEquals(t, errs, "teststruct.alphanumspace", "AlphanumSpace darf nur alphanumerische Zeichen und Leerzeichen enthalten")
	containsEquals(t, errs, "teststruct.alphaunicode", "AlphaUnicode darf nur Unicode-Buchstaben enthalten")
	containsEquals(t, errs, "teststruct.alphanumunicode", "AlphanumUnicode darf nur Unicode-alphanumerische Zeichen enthalten")
	containsEquals(t, errs, "teststruct.urnrfc2141", "UrnRFC2141 muss eine gültige RFC 2141 URN sein")
	containsEquals(t, errs, "teststruct.timezone", "Timezone muss eine gültige Zeitzone sein")
	containsEquals(t, errs, "teststruct.bcp47languagetag", "Bcp47LanguageTag muss ein gültiger BCP 47-Sprachcode sein")
	containsEquals(t, errs, "teststruct.validatefn", "ValidateFn muss ein gültiges Objekt sein")
}
