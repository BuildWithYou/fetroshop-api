package validatorhelper

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func GetValidator() *validator.Validate {
	return validator.New()
}

func IsNil(value any) bool {
	return value == nil
}

func IsNotNil(value any) bool {
	return value != nil
}

func IsZero(value int64) bool {
	return value == int64(0)
}

func IsNotZero(value int64) bool {
	return value != int64(0)
}

func ValidateBodyPayload(ctx *fiber.Ctx, vld *validator.Validate, payload any) (err error) {
	err = ctx.BodyParser(payload)
	if IsNotNil(err) {
		return err
	}

	err = vld.Struct(payload)
	if IsNotNil(err) {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return err
}

func ValidateQueryPayload(ctx *fiber.Ctx, vld *validator.Validate, payload any) (err error) {
	err = ctx.QueryParser(payload)
	if IsNotNil(err) {
		return err
	}

	err = vld.Struct(payload)
	if IsNotNil(err) {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return err
}

func ValidateParamPayload(ctx *fiber.Ctx, vld *validator.Validate, payload any) (err error) {
	err = ctx.ParamsParser(payload)
	if IsNotNil(err) {
		return err
	}

	err = vld.Struct(payload)
	if IsNotNil(err) {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return err
}

func ValidateCookiePayload(ctx *fiber.Ctx, vld *validator.Validate, payload any) (err error) {
	err = ctx.CookieParser(payload)
	if IsNotNil(err) {
		return err
	}

	err = vld.Struct(payload)
	if IsNotNil(err) {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return err
}
