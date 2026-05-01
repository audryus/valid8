package valid8

import (
	locale "github.com/go-playground/locales/es"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	translations "github.com/go-playground/validator/v10/translations/es"
)

func registerES(v *Valid8) {
	locale := locale.New()
	trans := ut.New(locale)
	translator, _ := trans.GetTranslator("es")
	translations.RegisterDefaultTranslations(v.Validator, translator)
	registerCustomESTranslations(v, translator)
	v.locales["es"] = translator
}

func registerCustomESTranslations(v *Valid8, trans ut.Translator) error {
	translations := []customTranslation{
		{
			tag:         "required_unless",
			translation: "{0} es un campo obligatorio",
			override:    false,
		},
		{
			tag:         "required_with",
			translation: "{0} es un campo obligatorio",
			override:    false,
		},
		{
			tag:         "required_with_all",
			translation: "{0} es un campo obligatorio",
			override:    false,
		},
		{
			tag:         "required_without",
			translation: "{0} es un campo obligatorio",
			override:    false,
		},
		{
			tag:         "required_without_all",
			translation: "{0} es un campo obligatorio",
			override:    false,
		},
		{
			tag:         "excluded_if",
			translation: "{0} es un campo excluido",
			override:    false,
		},
		{
			tag:         "excluded_unless",
			translation: "{0} es un campo excluido",
			override:    false,
		},
		{
			tag:         "excluded_with",
			translation: "{0} es un campo excluido",
			override:    false,
		},
		{
			tag:         "excluded_with_all",
			translation: "{0} es un campo excluido",
			override:    false,
		},
		{
			tag:         "excluded_without",
			translation: "{0} es un campo excluido",
			override:    false,
		},
		{
			tag:         "excluded_without_all",
			translation: "{0} es un campo excluido",
			override:    false,
		},
		{
			tag:         "isdefault",
			translation: "{0} debe estar en el valor predeterminado",
			override:    false,
		},
		{
			tag:         "alphaspace",
			translation: "{0} debe contener solo caracteres alfabéticos y espacios",
			override:    false,
		},
		{
			tag:         "alphanumspace",
			translation: "{0} debe contener solo caracteres alfanuméricos y espacios",
			override:    false,
		},
		{
			tag:         "alphaunicode",
			translation: "{0} debe contener solo caracteres unicode alfabéticos",
			override:    false,
		},
		{
			tag:         "alphanumunicode",
			translation: "{0} debe contener solo caracteres unicode alfanuméricos",
			override:    false,
		},
		{
			tag:         "urn_rfc2141",
			translation: "{0} debe ser un URN RFC 2141 válido",
			override:    false,
		},
		{
			tag:         "fqdn",
			translation: "{0} debe ser un FQDN válido",
			override:    false,
		},
		{
			tag:         "cron",
			translation: "{0} debe ser una expresión cron válida",
			override:    false,
		},
		{
			tag:         "json",
			translation: "{0} debe ser una cadena json válida",
			override:    false,
		},
		{
			tag:         "jwt",
			translation: "{0} debe ser una cadena jwt válida",
			override:    false,
		},
		{
			tag:         "lowercase",
			translation: "{0} debe ser una cadena en minúsculas",
			override:    false,
		},
		{
			tag:         "uppercase",
			translation: "{0} debe ser una cadena en mayúsculas",
			override:    false,
		},
		{
			tag:         "datetime",
			translation: "{0} no coincide con el formato {1}",
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
			translation: "{0} debe ser una zona horaria válida",
			override:    false,
		},
		{
			tag:         "postcode_iso3166_alpha2",
			translation: "{0} no coincide con el formato de código postal del país {1}",
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
			translation: "{0} no coincide con el formato de código postal del campo {1}",
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
			tag:         "bcp47_strict_language_tag",
			translation: "{0} debe ser un código de idioma BCP 47 válido",
			override:    false,
		},
	}

	return registerTranslations(v, trans, translations)
}
