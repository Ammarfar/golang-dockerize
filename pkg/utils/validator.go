package utils

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

func Validate(s interface{}) error {
	validate := validator.New(validator.WithRequiredStructEnabled())

	var ve validator.ValidationErrors

	if err := validate.Struct(s); err != nil {
		if errors.As(err, &ve) {
			for _, fe := range ve {
				return errors.New(msgForTag(fe))
			}
		}
	}

	return nil
}

func msgForTag(fe validator.FieldError) string {
	switch fe.Tag() {

	case "required":
		return fmt.Sprintf("%s is required", fe.Field())
	case "email":
		return "Invalid email"

	}

	return fe.Error() // default error
}
