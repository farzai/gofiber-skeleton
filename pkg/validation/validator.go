package validation

import (
	"strings"

	"github.com/farzai/app/pkg/support/str"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	uni      *ut.UniversalTranslator
)

func Validate(data interface{}) error {
	validate := validator.New()

	errs := validate.Struct(data)

	en := en.New()
	uni = ut.New(en, en)
	trans, _ := uni.GetTranslator("en")

	en_translations.RegisterDefaultTranslations(validate, trans)

	if errs != nil {
		bag := NewMessageBag()

		// Loop through errors
		for _, err := range errs.(validator.ValidationErrors) {
			errorMessage := err.Translate(trans)

			// Replace field name with snake case
			// e.g. PhoneNumber -> phone_number
			errorMessage = strings.Replace(errorMessage, err.Field(), str.SnakeCase(err.Field()), 1)

			// Add error to message bag
			bag.Add(str.SnakeCase(err.Field()), errorMessage)
		}

		return NewValidationError(errs, bag)
	}

	return nil
}
