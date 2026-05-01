package valid8

import (
	locale "github.com/go-playground/locales/fr"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	translations "github.com/go-playground/validator/v10/translations/fr"
)

func registerFR(v *Valid8) {
	locale := locale.New()
	trans := ut.New(locale)
	translator, _ := trans.GetTranslator("fr")
	translations.RegisterDefaultTranslations(v.Validator, translator)
	registerCustomFRTranslations(v, translator)
	v.locales["fr"] = translator
}

func registerCustomFRTranslations(v *Valid8, trans ut.Translator) error {
	translations := []customTranslation{
		{
			tag:         "required_if",
			translation: "{0} est un champ requis",
			override:    false,
		},
		{
			tag:         "required_unless",
			translation: "{0} est un champ requis",
			override:    false,
		},
		{
			tag:         "required_with",
			translation: "{0} est un champ requis",
			override:    false,
		},
		{
			tag:         "required_with_all",
			translation: "{0} est un champ requis",
			override:    false,
		},
		{
			tag:         "required_without",
			translation: "{0} est un champ requis",
			override:    false,
		},
		{
			tag:         "required_without_all",
			translation: "{0} est un champ requis",
			override:    false,
		},
		{
			tag:         "excluded_if",
			translation: "{0} est un champ exclu",
			override:    false,
		},
		{
			tag:         "excluded_unless",
			translation: "{0} est un champ exclu",
			override:    false,
		},
		{
			tag:         "excluded_with",
			translation: "{0} est un champ exclu",
			override:    false,
		},
		{
			tag:         "excluded_with_all",
			translation: "{0} est un champ exclu",
			override:    false,
		},
		{
			tag:         "excluded_without",
			translation: "{0} est un champ exclu",
			override:    false,
		},
		{
			tag:         "excluded_without_all",
			translation: "{0} est un champ exclu",
			override:    false,
		},
		{
			tag:         "isdefault",
			translation: "{0} doit être dans la valeur par défaut",
			override:    false,
		},
		{
			tag:         "alphaspace",
			translation: "{0} doit contenir uniquement des caractères alphabétiques et des espaces",
			override:    false,
		},
		{
			tag:         "alphanumspace",
			translation: "{0} doit contenir uniquement des caractères alphanumériques et des espaces",
			override:    false,
		},
		{
			tag:         "alphaunicode",
			translation: "{0} doit contenir uniquement des caractères unicode alphabétiques",
			override:    false,
		},
		{
			tag:         "alphanumunicode",
			translation: "{0} doit contenir uniquement des caractères unicode alphanumériques",
			override:    false,
		},
		{
			tag:         "e164",
			translation: "{0} doit être un numéro de téléphone au format E.164 valide",
			override:    false,
		},
		{
			tag:         "urn_rfc2141",
			translation: "{0} doit être un URN RFC 2141 valide",
			override:    false,
		},
		{
			tag:         "fqdn",
			translation: "{0} doit être un FQDN valide",
			override:    false,
		},
		{
			tag:         "unique",
			translation: "{0} doit contenir des valeurs uniques",
			override:    false,
		},
		{
			tag:         "cron",
			translation: "{0} doit être une expression cron valide",
			override:    false,
		},
		{
			tag:         "json",
			translation: "{0} doit être une chaîne json valide",
			override:    false,
		},
		{
			tag:         "jwt",
			translation: "{0} doit être une chaîne jwt valide",
			override:    false,
		},
		{
			tag:         "lowercase",
			translation: "{0} doit être une chaîne en minuscules",
			override:    false,
		},
		{
			tag:         "uppercase",
			translation: "{0} doit être une chaîne en majuscules",
			override:    false,
		},
		{
			tag:         "datetime",
			translation: "{0} ne correspond pas au format {1}",
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
			translation: "{0} doit être un fuseau horaire valide",
			override:    false,
		},
		{
			tag:         "postcode_iso3166_alpha2",
			translation: "{0} ne correspond pas au format de code postal du pays {1}",
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
			translation: "{0} ne correspond pas au format de code postal du champ {1}",
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
			translation: "{0} doit être un code de langue BCP 47 valide",
			override:    false,
		},
	}

	return registerTranslations(v, trans, translations)
}
