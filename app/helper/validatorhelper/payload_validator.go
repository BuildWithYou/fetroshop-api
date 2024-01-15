package validatorhelper

import (
	"fmt"

	"github.com/BuildWithYou/fetroshop-api/app/helper/errorhelper"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// TODO: handle proper message here

func generateErrorMessage(err error) error {
	// make error map
	errorMap := make(map[string]string)
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

	return errorhelper.Error400(errorMap)

}

func ValidateBodyPayload(ctx *fiber.Ctx, vld *validator.Validate, payload interface{}) (err error) {
	err = ctx.BodyParser(payload)
	if err != nil {
		return err
	}

	err = vld.Struct(payload)
	if err != nil {
		return generateErrorMessage(err)
	}

	return err
}

func ValidateQueryPayload(ctx *fiber.Ctx, vld *validator.Validate, payload interface{}) (err error) {
	err = ctx.QueryParser(payload)
	if err != nil {
		return err
	}

	err = vld.Struct(payload)
	if err != nil {
		return generateErrorMessage(err)
	}

	return err
}

func ValidateParamPayload(ctx *fiber.Ctx, vld *validator.Validate, payload interface{}) (err error) {
	err = ctx.ParamsParser(payload)
	if err != nil {
		return err
	}

	err = vld.Struct(payload)
	if err != nil {
		return generateErrorMessage(err)
	}

	return err
}

func ValidateCookiePayload(ctx *fiber.Ctx, vld *validator.Validate, payload interface{}) (err error) {
	err = ctx.CookieParser(payload)
	if err != nil {
		return err
	}

	err = vld.Struct(payload)
	if err != nil {
		return generateErrorMessage(err)
	}

	return err
}
