package validation

import "github.com/go-playground/validator/v10"

type Validation struct {
	validator *validator.Validate
}

func NewValidation() *Validation {
	validator := validator.New()
	return &Validation{validator}
}

func (v *Validation) Validate(i interface{}) error {
	return v.validator.Struct(i)
}
