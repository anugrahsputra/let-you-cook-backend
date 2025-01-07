package validator

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// ValidateStruct is a helper function to validate struct and return error messages.
func ValidateStruct(s interface{}) map[string]string {
	err := validate.Struct(s)
	if err == nil {
		return nil
	}

	errors := make(map[string]string)
	for _, err := range err.(validator.ValidationErrors) {
		errors[err.Field()] = formatErrorMessage(err)
	}

	return errors
}

func formatErrorMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email format"
	case "min":
		return "Minimum length is " + err.Param()
	case "max":
		return "Maximum length is " + err.Param()
	default:
		return "Invalid value"
	}
}
