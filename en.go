package valid8

import (
	locale "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	translations "github.com/go-playground/validator/v10/translations/en"
)

func registerEN(v *Valid8) {
	locale := locale.New()
	trans := ut.New(locale)
	translator, _ := trans.GetTranslator("en")
	translations.RegisterDefaultTranslations(v.Validator, translator)
	v.locales["en"] = translator
}
