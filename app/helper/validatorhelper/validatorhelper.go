package validatorhelper

import (
	"github.com/go-playground/validator/v10"
)

func GetValidator() *validator.Validate {
	return validator.New()
}

func IsZero(value int64) bool {
	return value == int64(0)
}

func IsNotZero(value int64) bool {
	return value != int64(0)
}
