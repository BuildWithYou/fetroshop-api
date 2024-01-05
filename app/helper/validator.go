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

func IsZero64(value int64) bool {
	return value == 0
}

func IsNotZero64(value int64) bool {
	return value != 0
}
