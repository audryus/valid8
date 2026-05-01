package valid8

import (
	locale "github.com/go-playground/locales/de"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	translations "github.com/go-playground/validator/v10/translations/de"
)

func registerZHTW(v *Valid8) {
	locale := locale.New()
	trans := ut.New(locale)
	translator, _ := trans.GetTranslator("zh_TW")
	translations.RegisterDefaultTranslations(v.Validator, translator)
	registerCustomZHTWTranslations(v, translator)
	v.locales["zh_TW"] = translator
}

func registerCustomZHTWTranslations(v *Valid8, trans ut.Translator) error {
	translations := []customTranslation{
		{
			tag:         "alphaspace",
			translation: "{0}只能包含字母和空格",
			override:    false,
		},
		{
			tag:         "alphanumspace",
			translation: "{0}只能包含字母、數字和空格",
			override:    false,
		},
		{
			tag:         "alphaunicode",
			translation: "{0}只能包含Unicode字母",
			override:    false,
		},
		{
			tag:         "alphanumunicode",
			translation: "{0}只能包含Unicode字母和數字",
			override:    false,
		},
		{
			tag:         "bcp47_strict_language_tag",
			translation: "{0}必須是有效的BCP 47語言標籤",
			override:    false,
		},
		{
			tag:         "cron",
			translation: "{0}必須是有效的cron表示式",
			override:    false,
		},
		{
			tag:         "fqdn",
			translation: "{0}必須是有效的FQDN",
			override:    false,
		},
		{
			tag:         "json",
			translation: "{0}必須是有效的JSON字串",
			override:    false,
		},
		{
			tag:         "jwt",
			translation: "{0}必須是有效的JWT字串",
			override:    false,
		},
		{
			tag:         "postcode_iso3166_alpha2",
			translation: "{0}不符合{1}國家的郵遞區號格式",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					return logErrorTranslation(fe)
				}
				return t
			},
		},
		{
			tag:         "postcode_iso3166_alpha2_field",
			translation: "{0}不符合欄位{1}指定國家的郵遞區號格式",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					return logErrorTranslation(fe)
				}
				return t
			},
		},
		{
			tag:         "timezone",
			translation: "{0}必須是有效的時區",
			override:    false,
		},
		{
			tag:         "unique",
			translation: "{0}裡的值必須是唯一的",
			override:    false,
		},
		{
			tag:         "uppercase",
			translation: "{0}必須是大寫字串",
			override:    false,
		},
		{
			tag:         "urn_rfc2141",
			translation: "{0}必須是有效的 RFC 2141 URN",
			override:    false,
		},
		{
			tag:         "validateFn",
			translation: "{0}必須是有效的物件",
			override:    false,
		},
	}

	return registerTranslations(v, trans, translations)
}
