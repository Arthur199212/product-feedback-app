package validation

import "github.com/go-playground/validator/v10"

type Validation struct {
	validator *validator.Validate
}

func NewValidation() *Validation {
	validator := validator.New()
	return &Validation{validator}
}

func (v *Validation) ValidateStruct(i interface{}) error {
	return v.validator.Struct(i)
}

func (v *Validation) ValidateVar(i interface{}, tag string) error {
	return v.validator.Var(i, tag)
}
