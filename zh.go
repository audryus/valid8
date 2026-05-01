package valid8

import (
	locale "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	translations "github.com/go-playground/validator/v10/translations/zh"
)

func registerZH(v *Valid8) {
	locale := locale.New()
	trans := ut.New(locale)
	translator, _ := trans.GetTranslator("zh")
	translations.RegisterDefaultTranslations(v.Validator, translator)
	registerCustomZHTranslations(v, translator)
	v.locales["zh"] = translator
}

func registerCustomZHTranslations(v *Valid8, trans ut.Translator) error {
	translations := []customTranslation{
		{
			tag:         "alphaspace",
			translation: "{0}只能包含字母和空格",
			override:    false,
		},
		{
			tag:         "bcp47_strict_language_tag",
			translation: "{0}必须是有效的BCP 47语言标签",
			override:    false,
		},
		{
			tag:         "cron",
			translation: "{0}必须是有效的cron表达式",
			override:    false,
		},
		{
			tag:         "fqdn",
			translation: "{0}必须是有效的FQDN",
			override:    false,
		},
		{
			tag:         "jwt",
			translation: "{0}必须是有效的JWT字符串",
			override:    false,
		},
		{
			tag:         "postcode_iso3166_alpha2",
			translation: "{0}不符合{1}国家的邮政编码格式",
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
			translation: "{0}不符合{1}字段所在国家的邮政编码格式",
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
			translation: "{0}必须是有效的时区",
			override:    false,
		},
		{
			tag:         "unique",
			translation: "{0}中的值必须唯一",
			override:    false,
		},
		{
			tag:         "urn_rfc2141",
			translation: "{0}必须是有效的 RFC 2141 URN",
			override:    false,
		},
		{
			tag:         "validateFn",
			translation: "{0}必须是一个有效对象",
			override:    false,
		},
	}

	return registerTranslations(v, trans, translations)
}
