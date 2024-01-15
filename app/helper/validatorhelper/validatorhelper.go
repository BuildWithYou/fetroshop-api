package validatorhelper

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

func IsZero(value int64) bool {
	return value == int64(0)
}

func IsNotZero(value int64) bool {
	return value != int64(0)
}
