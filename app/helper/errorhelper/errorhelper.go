package errorhelper

import (
	"github.com/gofiber/fiber/v2"
)

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func Error400(msg string) error {
	return fiber.NewError(fiber.StatusBadRequest, msg)
}

func Error500(msg string) error {
	return fiber.NewError(fiber.StatusInternalServerError, msg)
}

func ErrorCustom(errorCode int, msg string) error {
	return fiber.NewError(errorCode, msg)
}
