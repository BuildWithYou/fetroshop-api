package validatorhelper

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func generateErrorMessage(err error) (errValidation fiber.Map) {
	// make error map
	errValidation = make(fiber.Map)
	validationErrors := err.(validator.ValidationErrors)
	for _, fieldError := range validationErrors {
		fieldName := fieldError.Field()
		switch fieldError.Tag() {
		case "required":
			{
				errValidation[fieldName] = fmt.Sprint(fieldName, " field is required")
			}
		case "email":
			{
				errValidation[fieldName] = fmt.Sprint(fieldName, " field must be valid email format")
			}
		case "min":
			{
				if fieldError.Kind() == reflect.String {
					errValidation[fieldName] = fmt.Sprint(fieldName, " field must be longer than ", fieldError.Param(), " characters")
				} else {
					errValidation[fieldName] = fmt.Sprint(fieldName, " field must be greater than ", fieldError.Param())
				}
			}
		case "required_with":
			{
				fieldParam := fieldError.Param()
				fieldSlice := strings.Split(fieldParam, " ")
				for i, field := range fieldSlice {

					// Convert the first character to lowercase
					firstCharLower := strings.ToLower(string(field[0]))

					// Convert the last character to lowercase
					lastCharLower := strings.ToLower(string(field[len(field)-1]))

					// Combine the modified first and last characters with the rest of the string
					fieldSlice[i] = firstCharLower + field[1:len(field)-1] + lastCharLower
				}
				errValidation[fieldName] = fmt.Sprint(fieldName, " field is required when ", strings.Join(fieldSlice, ", "), " is filled")
			}
		default:
			{
				errValidation[fieldName] = fmt.Sprint("Error on tag ", fieldError.Tag(), " on field ", fieldName, " with error ", fieldError.Error())
			}
		}

	}

	return errValidation

}

func ValidateBodyPayload(ctx *fiber.Ctx, vld *validator.Validate, payload interface{}) (errValidation fiber.Map, errParsing error) {
	errParsing = ctx.BodyParser(payload)
	if errParsing != nil {
		return nil, errParsing
	}

	err := vld.Struct(payload)
	if err != nil {
		errValidation := generateErrorMessage(err)
		return errValidation, nil
	}

	return nil, nil
}

func ValidateQueryPayload(ctx *fiber.Ctx, vld *validator.Validate, payload interface{}) (errValidation fiber.Map, errParsing error) {
	errParsing = ctx.QueryParser(payload)
	if errParsing != nil {
		return nil, errParsing
	}

	err := vld.Struct(payload)
	if err != nil {
		errValidation := generateErrorMessage(err)
		return errValidation, nil
	}

	return nil, nil
}

func ValidateParamPayload(ctx *fiber.Ctx, vld *validator.Validate, payload interface{}) (errValidation fiber.Map, errParsing error) {
	errParsing = ctx.ParamsParser(payload)
	if errParsing != nil {
		return nil, errParsing
	}

	err := vld.Struct(payload)
	if err != nil {
		errValidation := generateErrorMessage(err)
		return errValidation, nil
	}

	return nil, nil
}

func ValidateCookiePayload(ctx *fiber.Ctx, vld *validator.Validate, payload interface{}) (errValidation fiber.Map, errParsing error) {
	errParsing = ctx.CookieParser(payload)
	if errParsing != nil {
		return nil, errParsing
	}

	err := vld.Struct(payload)
	if err != nil {
		errValidation := generateErrorMessage(err)
		return errValidation, nil
	}

	return nil, nil
}
