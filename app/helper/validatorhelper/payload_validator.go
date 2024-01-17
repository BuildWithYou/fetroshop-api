package validatorhelper

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// TODO: handle proper message here

func generateErrorMessage(err error) (errValidation fiber.Map) {
	// make error map
	errValidation = make(fiber.Map)
	validationErrors := err.(validator.ValidationErrors)
	for _, fieldError := range validationErrors {
		key := fieldError.Field()
		switch fieldError.Tag() {
		case "required":
			{
				errValidation[key] = fmt.Sprint(fieldError.Field(), " is required")
			}
		default:
			{
				errValidation[key] = fmt.Sprint("Error on tag ", fieldError.Tag(), " on field ", fieldError.Field(), " with error ", fieldError.Error())
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
