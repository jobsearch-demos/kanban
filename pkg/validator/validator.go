package validator

import (
	"github.com/creasty/defaults"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type WithDefaultsValidator struct {
	*validator.Validate
}

// NewWithDefaultsValidator creates a new validator for model fields.
func NewWithDefaultsValidator() WithDefaultsValidator {
	// Create a new validator for a Book model.
	validate := validator.New()
	// Custom validation for uuid.UUID fields.
	_ = validate.RegisterValidation("uuid", func(fl validator.FieldLevel) bool {
		field, ok := fl.Field().Interface().(uuid.UUID)

		if !ok {
			if _, err := uuid.Parse(fl.Field().String()); err != nil {
				return false
			}
		}
		value := field.String()

		if _, err := uuid.Parse(value); err != nil {
			return false
		}

		return true
	})

	return WithDefaultsValidator{
		validate,
	}
}

func (v *WithDefaultsValidator) ValidateWithDefaults(dto interface{}) error {
	// set default values
	if err := defaults.Set(dto); err != nil {
		return err
	}

	if err := v.Struct(dto); err != nil {
		return err
	}
	return nil
}

// ValidatorErrors func shows validation errors for each invalid fields.
func ValidatorErrors(err error) map[string]string {
	// Define fields map.
	fields := map[string]string{}

	// Make error message for each invalid field.
	for _, err := range err.(validator.ValidationErrors) {
		fields[err.Field()] = err.Error()
	}

	return fields
}
