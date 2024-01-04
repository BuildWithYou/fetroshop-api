package helper

import "github.com/go-playground/validator/v10"

func GetValidator() *validator.Validate {
	return validator.New()
}

func IsNil(value any) bool {
	return value == nil
}

func IsNotNil(value any) bool {
	return value != nil
}
