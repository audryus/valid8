package valid8_test

import (
	"testing"

	"codeberg.org/audryus/valid8"
)

func TestJATranslations(t *testing.T) {
	v := valid8.New(valid8.WithLocales(valid8.JA))
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
		AlphaSpace            string `validate:"alphaspace"`
		AlphanumSpace         string `validate:"alphanumspace"`
		AlphaUnicode          string `validate:"alphaunicode"`
		AlphanumUnicode       string `validate:"alphanumunicode"`
		Bcp47LanguageTag      string `validate:"bcp47_language_tag"`
		Cron                  string `validate:"cron"`
		Fqdn                  string `validate:"fqdn"`
		PostcodeIso3166Alpha2 string `validate:"postcode_iso3166_alpha2=JP"`
		UrnRFC2141            string `validate:"urn_rfc2141"`
		ValidateFn            string `validate:"validateFn"`

		// Helper fields
		OtherField          string
		AnotherField        string
		MissingField        string
		AnotherMissingField string
	}

	s := TestStruct{
		// Trigger required_...
		OtherField:   "value",
		AnotherField: "value",

		// Trigger excluded_...
		ExcludedIf:         "present",
		ExcludedUnless:     "present",
		ExcludedWith:       "present",
		ExcludedWithAll:    "present",
		ExcludedWithout:    "present",
		ExcludedWithoutAll: "present",

		// Trigger others with invalid data
		AlphaSpace:            "123",
		AlphanumSpace:         "!",
		AlphaUnicode:          "123",
		AlphanumUnicode:       "!",
		Bcp47LanguageTag:      "!!!",
		Cron:                  "abc",
		Fqdn:                  "abc",
		PostcodeIso3166Alpha2: "abc",
		UrnRFC2141:            "abc",
		ValidateFn:            "invalid",
	}

	err := v.Struct(s, valid8.JA)
	if len(err) == 0 {
		t.Fatal("Expected validation errors, got none")
	}

	errs := valid8.ErrorsToMap(err)

	// Required variants
	containsEquals(t, errs, "teststruct.requiredunless", "RequiredUnlessは必須フィールドです")
	containsEquals(t, errs, "teststruct.requiredwith", "RequiredWithは必須フィールドです")
	containsEquals(t, errs, "teststruct.requiredwithall", "RequiredWithAllは必須フィールドです")
	containsEquals(t, errs, "teststruct.requiredwithout", "RequiredWithoutは必須フィールドです")
	containsEquals(t, errs, "teststruct.requiredwithoutall", "RequiredWithoutAllは必須フィールドです")

	// Excluded variants
	containsEquals(t, errs, "teststruct.excludedif", "ExcludedIfは除外フィールドです")
	containsEquals(t, errs, "teststruct.excludedunless", "ExcludedUnlessは除外フィールドです")
	containsEquals(t, errs, "teststruct.excludedwith", "ExcludedWithは除外フィールドです")
	containsEquals(t, errs, "teststruct.excludedwithall", "ExcludedWithAllは除外フィールドです")
	containsEquals(t, errs, "teststruct.excludedwithout", "ExcludedWithoutは除外フィールドです")
	containsEquals(t, errs, "teststruct.excludedwithoutall", "ExcludedWithoutAllは除外フィールドです")

	// Other tags
	containsEquals(t, errs, "teststruct.alphaspace", "AlphaSpaceは半角英字と空白のみを含めることができます")
	containsEquals(t, errs, "teststruct.alphanumspace", "AlphanumSpaceは半角英数字と空白のみを含めることができます")
	containsEquals(t, errs, "teststruct.alphaunicode", "AlphaUnicodeはUnicodeの文字のみを含めることができます")
	containsEquals(t, errs, "teststruct.alphanumunicode", "AlphanumUnicodeはUnicodeの英数字のみを含めることができます")
	containsEquals(t, errs, "teststruct.bcp47languagetag", "Bcp47LanguageTagは有効なBCP 47言語タグでなければなりません")
	containsEquals(t, errs, "teststruct.cron", "Cronは有効なcron式でなければなりません")
	containsEquals(t, errs, "teststruct.fqdn", "Fqdnは有効なFQDNでなければなりません")
	containsEquals(t, errs, "teststruct.postcodeiso3166alpha2", "PostcodeIso3166Alpha2はJP国の郵便番号の書式と一致しません")
	containsEquals(t, errs, "teststruct.urnrfc2141", "UrnRFC2141は有効なRFC 2141 URNでなければなりません")
	containsEquals(t, errs, "teststruct.validatefn", "ValidateFnは有効なオブジェクトでなければなりません")
}
