package valid8

import (
	localePTBR "github.com/go-playground/locales/pt_BR"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	translationsBR "github.com/go-playground/validator/v10/translations/pt_BR"
)

func registerBR(v *Valid8) {
	locale := localePTBR.New()
	trans := ut.New(locale)
	translator, _ := trans.GetTranslator("pt_BR")
	translationsBR.RegisterDefaultTranslations(v.Validator, translator)
	registerCustomPTBRTranslations(v, translator)
	v.locales["pt_BR"] = translator
}

func registerCustomPTBRTranslations(v *Valid8, trans ut.Translator) error {
	translations := []customTranslation{
		{
			tag:         "required_if",
			translation: "{0} é um campo obrigatório",
			override:    false,
		},
		{
			tag:         "required_unless",
			translation: "{0} é um campo obrigatório",
			override:    false,
		},
		{
			tag:         "required_with",
			translation: "{0} é um campo obrigatório",
			override:    false,
		},
		{
			tag:         "required_with_all",
			translation: "{0} é um campo obrigatório",
			override:    false,
		},
		{
			tag:         "required_without",
			translation: "{0} é um campo obrigatório",
			override:    false,
		},
		{
			tag:         "required_without_all",
			translation: "{0} é um campo obrigatório",
			override:    false,
		},
		{
			tag:         "excluded_if",
			translation: "{0} é um campo excluído",
			override:    false,
		},
		{
			tag:         "excluded_unless",
			translation: "{0} é um campo excluído",
			override:    false,
		},
		{
			tag:         "excluded_with",
			translation: "{0} é um campo excluído",
			override:    false,
		},
		{
			tag:         "excluded_with_all",
			translation: "{0} é um campo excluído",
			override:    false,
		},
		{
			tag:         "excluded_without",
			translation: "{0} é um campo excluído",
			override:    false,
		},
		{
			tag:         "excluded_without_all",
			translation: "{0} é um campo excluído",
			override:    false,
		},
		{
			tag:         "isdefault",
			translation: "{0} deve estar no valor padrão",
			override:    false,
		},
		{
			tag:         "alphaspace",
			translation: "{0} deve conter apenas caracteres alfabéticos e espaços",
			override:    false,
		},
		{
			tag:         "alphanumspace",
			translation: "{0} deve conter apenas caracteres alfanuméricos e espaços",
			override:    false,
		},
		{
			tag:         "alphaunicode",
			translation: "{0} deve conter apenas caracteres unicode alfabéticos",
			override:    false,
		},
		{
			tag:         "alphanumunicode",
			translation: "{0} deve conter apenas caracteres unicode alfanuméricos",
			override:    false,
		},
		{
			tag:         "e164",
			translation: "{0} deve ser um número de telefone no formato E.164 válido",
			override:    false,
		},
		{
			tag:         "urn_rfc2141",
			translation: "{0} deve ser um URN RFC 2141 válido",
			override:    false,
		},
		{
			tag:         "fqdn",
			translation: "{0} deve ser um FQDN válido",
			override:    false,
		},
		{
			tag:         "unique",
			translation: "{0} deve conter valores únicos",
			override:    false,
		},
		{
			tag:         "cron",
			translation: "{0} deve ser uma expressão cron válida",
			override:    false,
		},
		{
			tag:         "json",
			translation: "{0} deve ser uma string json válida",
			override:    false,
		},
		{
			tag:         "jwt",
			translation: "{0} deve ser uma string jwt válida",
			override:    false,
		},
		{
			tag:         "lowercase",
			translation: "{0} deve ser uma string em minúsculo",
			override:    false,
		},
		{
			tag:         "uppercase",
			translation: "{0} deve ser uma string em maiúsculo",
			override:    false,
		},
		{
			tag:         "datetime",
			translation: "{0} não corresponde ao formato {1}",
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
			translation: "{0} deve ser um fuso horário válido",
			override:    false,
		},
		{
			tag:         "postcode_iso3166_alpha2",
			translation: "{0} não corresponde ao formato de CEP do país {1}",
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
			translation: "{0} não corresponde ao formato de CEP do campo {1}",
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
			translation: "{0} deve ser um código de idioma BCP 47 válido",
			override:    false,
		},
	}

	return registerTranslations(v, trans, translations)
}
