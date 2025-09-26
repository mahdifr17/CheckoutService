package utils

import (
	"github.com/go-playground/validator/v10"
)

var Validator *validator.Validate

func GetValidator() *validator.Validate {
	if Validator == nil {
		Validator = validator.New()
	}

	return Validator
}
