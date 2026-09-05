package valid8_test

import (
	"testing"

	"codeberg.org/audryus/valid8"
)

func TestZHTWTranslations(t *testing.T) {
	v := valid8.New(valid8.WithLocales(valid8.ZHTW))
	if v == nil {
		t.Fatal("New() returned nil")
	}

	type TestStruct struct {
		AlphaSpace             string `validate:"alphaspace"`
		AlphanumSpace          string `validate:"alphanumspace"`
		AlphaUnicode           string `validate:"alphaunicode"`
		AlphanumUnicode        string `validate:"alphanumunicode"`
		Bcp47LanguageTag       string `validate:"bcp47_language_tag"`
		Cron                   string `validate:"cron"`
		Fqdn                   string `validate:"fqdn"`
		Json                   string `validate:"json"`
		Jwt                    string `validate:"jwt"`
		PostcodeIso3166Alpha2  string `validate:"postcode_iso3166_alpha2=TW"`
		PostcodeIso3166Alpha2F string `validate:"postcode_iso3166_alpha2_field=CountryCode"`
		Timezone               string `validate:"timezone"`
		Unique                 []int  `validate:"unique"`
		Uppercase              string `validate:"uppercase"`
		UrnRFC2141             string `validate:"urn_rfc2141"`
		ValidateFn             string `validate:"validateFn"`

		// Helper fields
		CountryCode string
	}

	s := TestStruct{
		AlphaSpace:             "123",
		AlphanumSpace:          "!",
		AlphaUnicode:           "123",
		AlphanumUnicode:        "!",
		Bcp47LanguageTag:       "!!!",
		Cron:                   "abc",
		Fqdn:                   "abc",
		Json:                   "abc",
		Jwt:                    "abc",
		PostcodeIso3166Alpha2:  "abc",
		PostcodeIso3166Alpha2F: "abc",
		Timezone:               "abc",
		Unique:                 []int{1, 1},
		Uppercase:              "abc",
		UrnRFC2141:             "abc",
		ValidateFn:             "invalid",
		CountryCode:            "TW",
	}

	err := v.Struct(s, valid8.ZHTW)
	if len(err) == 0 {
		t.Fatal("Expected validation errors, got none")
	}

	errs := valid8.ErrorsToMap(err)

	containsEquals(t, errs, "teststruct.alphaspace", "AlphaSpace只能包含字母和空格")
	containsEquals(t, errs, "teststruct.alphanumspace", "AlphanumSpace只能包含字母、數字和空格")
	containsEquals(t, errs, "teststruct.alphaunicode", "AlphaUnicode只能包含Unicode字母")
	containsEquals(t, errs, "teststruct.alphanumunicode", "AlphanumUnicode只能包含Unicode字母和數字")
	containsEquals(t, errs, "teststruct.bcp47languagetag", "Bcp47LanguageTag必須是有效的BCP 47語言標籤")
	containsEquals(t, errs, "teststruct.cron", "Cron必須是有效的cron表示式")
	containsEquals(t, errs, "teststruct.fqdn", "Fqdn必須是有效的FQDN")
	containsEquals(t, errs, "teststruct.json", "Json必須是有效的JSON字串")
	containsEquals(t, errs, "teststruct.jwt", "Jwt必須是有效的JWT字串")
	containsEquals(t, errs, "teststruct.postcodeiso3166alpha2", "PostcodeIso3166Alpha2不符合TW國家的郵遞區號格式")
	containsEquals(t, errs, "teststruct.postcodeiso3166alpha2f", "PostcodeIso3166Alpha2F不符合欄位CountryCode指定國家的郵遞區號格式")
	containsEquals(t, errs, "teststruct.timezone", "Timezone必須是有效的時區")
	containsEquals(t, errs, "teststruct.unique", "Unique裡的值必須是唯一的")
	containsEquals(t, errs, "teststruct.uppercase", "Uppercase必須是大寫字串")
	containsEquals(t, errs, "teststruct.urnrfc2141", "UrnRFC2141必須是有效的 RFC 2141 URN")
	containsEquals(t, errs, "teststruct.validatefn", "ValidateFn必須是有效的物件")
}
