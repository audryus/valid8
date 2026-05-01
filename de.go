package valid8

import (
	locale "github.com/go-playground/locales/de"
	ut "github.com/go-playground/universal-translator"
	translations "github.com/go-playground/validator/v10/translations/de"
)

func registerDE(v *Valid8) {
	locale := locale.New()
	trans := ut.New(locale)
	translator, _ := trans.GetTranslator("de")
	translations.RegisterDefaultTranslations(v.Validator, translator)
	registerCustomDETranslations(v, translator)
	v.locales["de"] = translator
}

func registerCustomDETranslations(v *Valid8, trans ut.Translator) error {
	translations := []customTranslation{
		{
			tag:         "alphaspace",
			translation: "{0} darf nur Buchstaben und Leerzeichen enthalten",
			override:    false,
		},
		{
			tag:         "alphanumspace",
			translation: "{0} darf nur alphanumerische Zeichen und Leerzeichen enthalten",
			override:    false,
		},
		{
			tag:         "alphaunicode",
			translation: "{0} darf nur Unicode-Buchstaben enthalten",
			override:    false,
		},
		{
			tag:         "alphanumunicode",
			translation: "{0} darf nur Unicode-alphanumerische Zeichen enthalten",
			override:    false,
		},
		{
			tag:         "urn_rfc2141",
			translation: "{0} muss eine gültige RFC 2141 URN sein",
			override:    false,
		},
		{
			tag:         "timezone",
			translation: "{0} muss eine gültige Zeitzone sein",
			override:    false,
		},
		{
			tag:         "bcp47_language_tag",
			translation: "{0} muss ein gültiger BCP 47-Sprachcode sein",
			override:    false,
		},
		{
			tag:         "validateFn",
			translation: "{0} muss ein gültiges Objekt sein",
			override:    false,
		},
	}

	return registerTranslations(v, trans, translations)
}
