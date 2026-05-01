package valid8

import (
	"errors"
	"log"
	"strings"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

var (
	// ErrTagMandatory is returned when a translation is registered without a tag.
	ErrTagMandatory = errors.New("Tag is mandatory")
	// ErrTextMandatory is returned when a translation is registered without text.
	ErrTextMandatory = errors.New("Text is mandatory")
	// ErrLocaleMandatory is returned when a translation is registered without a locale.
	ErrLocaleMandatory = errors.New("Locale is mandatory")
)

// ValidationErrors wraps a validator.FieldError and its corresponding translated message.
type ValidationErrors struct {
	validator.FieldError
	message string
}

// Namespace returns the field namespace (e.g., "User.Email") in lowercase.
func (v *ValidationErrors) Namespace() string {
	return strings.ToLower(v.FieldError.Namespace())
}

// Message returns the translated, human-readable error message.
func (v *ValidationErrors) Message() string {
	return v.message
}

// Valid8 is the main structure for handling multi-language struct validation.
type Valid8 struct {
	// Validator is the underlying go-playground/validator instance used for validation.
	Validator *validator.Validate
	locales   map[Locale]ut.Translator
}

// Translation defines the configuration for a custom validation tag translation.
type Translation struct {
	// IgnoreFallback if true, prevents the translation from being registered for the default English (EN) locale
	// if the primary Locale is different.
	IgnoreFallback bool
	// Text is the localized error message template (e.g., "{0} is invalid").
	Text string
	// IgnoreOverride if true, prevents the translation from overriding an existing one for the same tag and locale.
	IgnoreOverride bool
	// Locale is the target language for this translation.
	Locale Locale
	// Tag is the validation tag identifier (e.g., "required", "email", "custom_tag").
	Tag string
	// TranslationFn is the function responsible for generating the localized error message.
	// This field is optional; if not provided, it defaults to an internal translation function
	// that retrieves the message from the universal-translator using the specified Tag.
	TranslationFn validator.TranslationFunc
}

// defaultTranslationFn returns a standard TranslationFunc that retrieves the translated
// message for a specific tag from the universal-translator. It automatically maps
// the field name to the {0} placeholder in the translation template.
func defaultTranslationFn(tag string) func(ut ut.Translator, fe validator.FieldError) string {
	return func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(tag, fe.Field())
		return t
	}
}

// RegisterTranslation registers a custom translation for a specific locale.
// By default, it also registers the translation for the English (EN) locale as a fallback,
// unless IgnoreFallback is set to true. This ensures that the custom validation tag
// has a meaningful message even if the requested language is not available.
func (v *Valid8) RegisterTranslation(translation Translation) error {
	if translation.Tag == "" {
		return ErrTagMandatory
	}

	if translation.Text == "" {
		return ErrTextMandatory
	}

	if translation.Locale == "" {
		return ErrLocaleMandatory
	}

	loc, ok := v.locales[translation.Locale]

	translationFn := translation.TranslationFn

	if translationFn == nil {
		translationFn = defaultTranslationFn(translation.Tag)
	}

	registerFn := func(ut ut.Translator) error {
		return ut.Add(translation.Tag,
			translation.Text,
			!translation.IgnoreOverride)
	}

	if ok {
		v.Validator.RegisterTranslation(translation.Tag, loc,
			registerFn,
			translationFn)
	}
	if !translation.IgnoreFallback {
		v.Validator.RegisterTranslation(translation.Tag, v.locales[EN],
			registerFn,
			translationFn)
	}

	return nil
}

// Struct validates the provided struct 's' using the specified 'language' for error messages.
// If the specified language is not registered, it defaults to English (EN).
func (v *Valid8) Struct(s any, language Locale) []ValidationErrors {
	var errors []ValidationErrors

	err := v.Validator.Struct(s)

	trans, ok := v.locales[language]
	if !ok {
		trans = v.locales[EN]
	}

	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			errors = append(errors, ValidationErrors{
				e,
				e.Translate(trans),
			})
		}
	}

	return errors
}

type configFunc func(*Valid8)

// Locale represents a language identifier used for localized validation messages.
type Locale string

const (
	// EN represents the English locale.
	EN Locale = "en"
	// PT_BR represents the Brazilian Portuguese locale.
	PT_BR Locale = "pt_BR"
	// ES represents the Spanish locale.
	ES Locale = "es"
	// FR represents the French locale.
	FR Locale = "fr"
	// DE represents the German locale.
	DE Locale = "de"
	// JA represents the Japanese locale.
	JA Locale = "ja"
	// ZH represents the Chinese (Simplified) locale.
	ZH Locale = "zh"
	// ZHTW represents the Chinese (Traditional) locale.
	ZHTW Locale = "zh_TW"
)

var languages = map[Locale]configFunc{
	PT_BR: registerBR,
	ES:    registerES,
	FR:    registerFR,
	DE:    registerDE,
	JA:    registerJA,
	ZH:    registerZH,
	ZHTW:  registerZHTW,
}

// WithLocales configures the Valid8 instance to support a specific set of locales.
// English (EN) is always registered by default.
func WithLocales(locales ...Locale) configFunc {
	return func(v *Valid8) {
		registerEN(v)

		for _, lng := range locales {
			if fn, ok := languages[lng]; ok {
				fn(v)
			}
		}
	}
}

// WithAllLocales configures the Valid8 instance to support all available locales.
func WithAllLocales() configFunc {
	return func(v *Valid8) {
		registerEN(v)
		for _, lng := range languages {
			lng(v)
		}
	}
}

// New creates and initializes a new Valid8 instance with the provided configuration options.
func New(opts ...configFunc) *Valid8 {
	v := &Valid8{
		Validator: validator.New(),
		locales:   make(map[Locale]ut.Translator),
	}

	WithLocales(EN)(v)

	for _, opt := range opts {
		opt(v)
	}

	return v
}

// ErrorsToMap converts a slice of ValidationErrors into a map where the key is the field namespace
// and the value is the human-readable error message.
func ErrorsToMap(errors []ValidationErrors) map[string]string {
	mep := make(map[string]string)

	for i := range errors {
		err := errors[i]
		mep[err.Namespace()] = err.Message()
	}

	return mep
}

func registrationFunc(tag string, translation string, override bool) validator.RegisterTranslationsFunc {
	return func(ut ut.Translator) (err error) {
		if err = ut.Add(tag, translation, override); err != nil {
			return
		}

		return
	}
}

func translateFunc(ut ut.Translator, fe validator.FieldError) string {
	t, err := ut.T(fe.Tag(), fe.Field())
	if err != nil {
		return logErrorTranslation(fe)
	}

	return t
}

type customTranslation struct {
	tag             string
	translation     string
	override        bool
	customRegisFunc validator.RegisterTranslationsFunc
	customTransFunc validator.TranslationFunc
}

func registerTranslations(v *Valid8,
	trans ut.Translator,
	translations []customTranslation) error {

	var err error

	for _, t := range translations {
		if t.customTransFunc != nil && t.customRegisFunc != nil {
			err = v.Validator.RegisterTranslation(t.tag, trans, t.customRegisFunc, t.customTransFunc)
		} else if t.customTransFunc != nil && t.customRegisFunc == nil {
			err = v.Validator.RegisterTranslation(t.tag, trans, registrationFunc(t.tag, t.translation, t.override), t.customTransFunc)
		} else if t.customTransFunc == nil && t.customRegisFunc != nil {
			err = v.Validator.RegisterTranslation(t.tag, trans, t.customRegisFunc, translateFunc)
		} else {
			err = v.Validator.RegisterTranslation(t.tag, trans, registrationFunc(t.tag, t.translation, t.override), translateFunc)
		}

		if err != nil {
			log.Printf("warning: error registering translation for tag %s: %#v", t.tag, err)
		}
	}

	return err
}

func logErrorTranslation(fe validator.FieldError) string {
	log.Printf("warning: error translating FieldError: %#v", fe)
	return fe.(error).Error()
}
