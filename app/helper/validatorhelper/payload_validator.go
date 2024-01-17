package validatorhelper

import (
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// TODO: handle proper message here

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
