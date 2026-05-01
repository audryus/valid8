package valid8

import (
	locale "github.com/go-playground/locales/ja"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	translations "github.com/go-playground/validator/v10/translations/ja"
)

func registerJA(v *Valid8) {
	locale := locale.New()
	trans := ut.New(locale)
	translator, _ := trans.GetTranslator("ja")
	translations.RegisterDefaultTranslations(v.Validator, translator)
	registerCustomJATranslations(v, translator)
	v.locales["ja"] = translator
}

func registerCustomJATranslations(v *Valid8, trans ut.Translator) error {
	translations := []customTranslation{
		{
			tag:         "required_unless",
			translation: "{0}は必須フィールドです",
			override:    false,
		},
		{
			tag:         "required_with",
			translation: "{0}は必須フィールドです",
			override:    false,
		},
		{
			tag:         "required_with_all",
			translation: "{0}は必須フィールドです",
			override:    false,
		},
		{
			tag:         "required_without",
			translation: "{0}は必須フィールドです",
			override:    false,
		},
		{
			tag:         "required_without_all",
			translation: "{0}は必須フィールドです",
			override:    false,
		},
		{
			tag:         "excluded_if",
			translation: "{0}は除外フィールドです",
			override:    false,
		},
		{
			tag:         "excluded_unless",
			translation: "{0}は除外フィールドです",
			override:    false,
		},
		{
			tag:         "excluded_with",
			translation: "{0}は除外フィールドです",
			override:    false,
		},
		{
			tag:         "excluded_with_all",
			translation: "{0}は除外フィールドです",
			override:    false,
		},
		{
			tag:         "excluded_without",
			translation: "{0}は除外フィールドです",
			override:    false,
		},
		{
			tag:         "excluded_without_all",
			translation: "{0}は除外フィールドです",
			override:    false,
		},
		{
			tag:         "alphaspace",
			translation: "{0}は半角英字と空白のみを含めることができます",
			override:    false,
		},
		{
			tag:         "alphanumspace",
			translation: "{0}は半角英数字と空白のみを含めることができます",
			override:    false,
		},
		{
			tag:         "alphaunicode",
			translation: "{0}はUnicodeの文字のみを含めることができます",
			override:    false,
		},
		{
			tag:         "alphanumunicode",
			translation: "{0}はUnicodeの英数字のみを含めることができます",
			override:    false,
		},
		{
			tag:         "bcp47_strict_language_tag",
			translation: "{0}は有効なBCP 47言語タグでなければなりません",
			override:    false,
		},
		{
			tag:         "cron",
			translation: "{0}は有効なcron式でなければなりません",
			override:    false,
		},
		{
			tag:         "fqdn",
			translation: "{0}は有効なFQDNでなければなりません",
			override:    false,
		},
		{
			tag:         "postcode_iso3166_alpha2",
			translation: "{0}は{1}国の郵便番号の書式と一致しません",
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
			tag:         "urn_rfc2141",
			translation: "{0}は有効なRFC 2141 URNでなければなりません",
			override:    false,
		},
		{
			tag:         "validateFn",
			translation: "{0}は有効なオブジェクトでなければなりません",
			override:    false,
		},
	}

	return registerTranslations(v, trans, translations)
}
