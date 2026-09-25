package valid8_test

import (
	"testing"

	"github.com/audryus/valid8"
)

func TestZHTranslations(t *testing.T) {
	v := valid8.New(valid8.WithLocales(valid8.ZH))
	if v == nil {
		t.Fatal("New() returned nil")
	}

	type TestStruct struct {
		AlphaSpace             string `validate:"alphaspace"`
		Bcp47LanguageTag       string `validate:"bcp47_language_tag"`
		Cron                   string `validate:"cron"`
		Fqdn                   string `validate:"fqdn"`
		Jwt                    string `validate:"jwt"`
		PostcodeIso3166Alpha2  string `validate:"postcode_iso3166_alpha2=CN"`
		PostcodeIso3166Alpha2F string `validate:"postcode_iso3166_alpha2_field=CountryCode"`
		Timezone               string `validate:"timezone"`
		Unique                 []int  `validate:"unique"`
		UrnRFC2141             string `validate:"urn_rfc2141"`
		ValidateFn             string `validate:"validateFn"`

		// Helper fields
		CountryCode string
	}

	s := TestStruct{
		AlphaSpace:             "123",
		Bcp47LanguageTag:       "!!!",
		Cron:                   "abc",
		Fqdn:                   "abc",
		Jwt:                    "abc",
		PostcodeIso3166Alpha2:  "abc",
		PostcodeIso3166Alpha2F: "abc",
		Timezone:               "abc",
		Unique:                 []int{1, 1},
		UrnRFC2141:             "abc",
		ValidateFn:             "invalid",
		CountryCode:            "CN",
	}

	err := v.Struct(s, valid8.ZH)
	if len(err) == 0 {
		t.Fatal("Expected validation errors, got none")
	}

	errs := valid8.ErrorsToMap(err)

	containsEquals(t, errs, "teststruct.alphaspace", "AlphaSpace只能包含字母和空格")
	containsEquals(t, errs, "teststruct.bcp47languagetag", "Bcp47LanguageTag必须是有效的BCP 47语言标签")
	containsEquals(t, errs, "teststruct.cron", "Cron必须是有效的cron表达式")
	containsEquals(t, errs, "teststruct.fqdn", "Fqdn必须是有效的FQDN")
	containsEquals(t, errs, "teststruct.jwt", "Jwt必须是有效的JWT字符串")
	containsEquals(t, errs, "teststruct.postcodeiso3166alpha2", "PostcodeIso3166Alpha2不符合CN国家的邮政编码格式")
	containsEquals(t, errs, "teststruct.postcodeiso3166alpha2f", "PostcodeIso3166Alpha2F不符合CountryCode字段所在国家的邮政编码格式")
	containsEquals(t, errs, "teststruct.timezone", "Timezone必须是有效的时区")
	containsEquals(t, errs, "teststruct.unique", "Unique中的值必须唯一")
	containsEquals(t, errs, "teststruct.urnrfc2141", "UrnRFC2141必须是有效的 RFC 2141 URN")
	containsEquals(t, errs, "teststruct.validatefn", "ValidateFn必须是一个有效对象")
}
