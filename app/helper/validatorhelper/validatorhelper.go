package validatorhelper

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var Validator *validator.Validate

func GetValidator() *validator.Validate {
	if Validator == nil {
		Validator = validator.New()
		Validator.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

			if name == "-" {
				return ""
			}

			return name
		})

	}
	return Validator
}

func IsZero(value int64) bool {
	return value == int64(0)
}

func IsNotZero(value int64) bool {
	return value != int64(0)
}

func IsValidUrl(urlInput string) bool {
	uri, err := url.Parse(urlInput)

	fmt.Println(uri)
	if err != nil {
		return false
	}

	if uri.Scheme != "http" && uri.Scheme != "https" {
		return false
	}

	return true
}
