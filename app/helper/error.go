package helper

import (
	"errors"

	"github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func ErrorIsNil(err error) bool {
	return err == nil
}

func ErrorIsNotNil(err error) bool {
	return err != nil
}

func Error500(ctx *fiber.Ctx, err error) error {
	return ctx.Status(fiber.StatusInternalServerError).JSON(model.Response{
		Code:    fiber.ErrInternalServerError.Code,
		Status:  fiber.ErrInternalServerError.Message,
		Message: err.Error(),
	})
}

func ErrorCustom(ctx *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError
	status := fiber.ErrInternalServerError.Message

	// Retrieve the custom status code if it's a *fiber.Error
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
		status = utils.StatusMessage(e.Code)
	}

	return ctx.Status(code).JSON(model.Response{
		Code:    code,
		Status:  status,
		Message: err.Error(),
	})
}
