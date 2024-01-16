package validatorhelper

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// TODO: handle proper message here

func generateErrorMessage(err error) fiber.Map {
	// make error map
	errorMap := make(fiber.Map)
	validationErrors := err.(validator.ValidationErrors)
	for _, fieldError := range validationErrors {
		key := fieldError.Field()
		switch fieldError.Tag() {
		case "required":
			{
				errorMap[key] = fmt.Sprint(fieldError.Field(), " is required")
			}
		default:
			{
				errorMap[key] = fmt.Sprint("Error on tag ", fieldError.Tag(), " on field ", fieldError.Field(), " with error ", fieldError.Error())
			}
		}

	}

	return errorMap

}

func ValidateBodyPayload(ctx *fiber.Ctx, vld *validator.Validate, payload interface{}) (errorMap fiber.Map, err error) {
	err = ctx.BodyParser(payload)
	if err != nil {
		return nil, err
	}

	err = vld.Struct(payload)
	if err != nil {
		errorMap := generateErrorMessage(err)
		return errorMap, nil
	}

	return nil, nil
}

func ValidateQueryPayload(ctx *fiber.Ctx, vld *validator.Validate, payload interface{}) (errorMap fiber.Map, err error) {
	err = ctx.QueryParser(payload)
	if err != nil {
		return nil, err
	}

	err = vld.Struct(payload)
	if err != nil {
		errorMap := generateErrorMessage(err)
		return errorMap, nil
	}

	return nil, nil
}

func ValidateParamPayload(ctx *fiber.Ctx, vld *validator.Validate, payload interface{}) (errorMap fiber.Map, err error) {
	err = ctx.ParamsParser(payload)
	if err != nil {
		return nil, err
	}

	err = vld.Struct(payload)
	if err != nil {
		errorMap := generateErrorMessage(err)
		return errorMap, nil
	}

	return nil, nil
}

func ValidateCookiePayload(ctx *fiber.Ctx, vld *validator.Validate, payload interface{}) (errorMap fiber.Map, err error) {
	err = ctx.CookieParser(payload)
	if err != nil {
		return nil, err
	}

	err = vld.Struct(payload)
	if err != nil {
		errorMap := generateErrorMessage(err)
		return errorMap, nil
	}

	return nil, nil
}
