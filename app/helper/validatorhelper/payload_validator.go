package validatorhelper

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// TODO: handle proper message here

func ValidateBodyPayload(ctx *fiber.Ctx, vld *validator.Validate, payload interface{}) (err error) {
	err = ctx.BodyParser(payload)
	if err != nil {
		return err
	}

	err = vld.Struct(payload)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
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
		// return fiber.NewError(fiber.StatusBadRequest, err.Error())

		validationErrors := err.(validator.ValidationErrors)
		/* for _,fieldError := range validationErrors {

		   } */
		errorMesage := fmt.Sprint("Erron on tag ", validationErrors[0].Tag(), " with error ", validationErrors[0].Error())
		return fiber.NewError(fiber.StatusBadRequest, errorMesage)
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
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
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
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return err
}
