package utility

import (
	"activitypub"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// Validate validates the input struct
func Validate(payload interface{}) error {
	err := validate.Struct(payload)

	if err != nil {
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(
				validationErrors,
				fmt.Sprintf("`%v` with value `%v` doesn't satisfy the `%v` constraint", err.Field(), err.Value(), err.Tag()),
			)
		}

		return errors.New(strings.Join(validationErrors, ","))
	}

	return nil
}

// CUSTOM VALIDATION RULES =============================================

var usernameRegex *regexp.Regexp

func init() {
	usernameRegex, _ = regexp.Compile(`^[a-z0-9]+(?:-[a-z0-9]+)*$`)
}

var _ = validate.RegisterValidation("username", func(fl validator.FieldLevel) bool {
	return usernameRegex.MatchString(fl.Field().String())
})

// Password validation rule: required,min=6,max=100
var _ = validate.RegisterValidation("password", func(fl validator.FieldLevel) bool {
	l := len(fl.Field().String())

	return l >= 6 && l < 100
})

var _ = validate.RegisterValidation("activitystream", func(fl validator.FieldLevel) bool {
	return fl.Field().String() == activitypub.ActivityStreams
})
