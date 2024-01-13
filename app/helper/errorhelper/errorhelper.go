package errorhelper

import (
	"fmt"

	"github.com/BuildWithYou/fetroshop-api/app/helper/logger"
	"github.com/gofiber/fiber/v2"
)

func PanicIfError(err error) {
	if err != nil {
		logger := logger.NewFrameworkLogger()
		logger.Panic(fmt.Sprint("Error : ", err.Error()))
	}
}

func Error400(msg string) error {
	return fiber.NewError(fiber.StatusBadRequest, msg)
}

func Error401(msg string) error {
	return fiber.NewError(fiber.StatusUnauthorized, msg)
}

func Error500(msg string) error {
	return fiber.NewError(fiber.StatusInternalServerError, msg)
}

func ErrorCustom(errorCode int, msg string) error {
	return fiber.NewError(errorCode, msg)
}
