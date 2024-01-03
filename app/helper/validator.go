package helper

import "github.com/go-playground/validator/v10"

func GetValidator() *validator.Validate {
	return validator.New()
}
