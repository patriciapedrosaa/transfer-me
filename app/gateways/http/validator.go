package http

import "github.com/go-playground/validator/v10"

type JSONValidator struct {
	validate *validator.Validate
}

func NewJSONValidator() JSONValidator {
	validate := validator.New()
	return JSONValidator{
		validate,
	}
}

func (j JSONValidator) Validate(data interface{}) error {
	err := j.validate.Struct(data)
	return err
}
