package errorhelper

import (
	"fmt"

	"github.com/BuildWithYou/fetroshop-api/app/helper/logger"
	"github.com/gofiber/fiber/v2"
)

var fwLogger = logger.NewFrameworkLogger()

func Error400(msg string) error {
	return fiber.NewError(fiber.StatusBadRequest, msg)
}

func Error401(msg string) error {
	return fiber.NewError(fiber.StatusUnauthorized, msg)
}

func Error500(msg string) error {
	fwLogger.Error(fmt.Sprint("Error : ", msg))
	return fiber.NewError(fiber.StatusInternalServerError, msg)
}

func ErrorCustom(errorCode int, msg string) error {
	return fiber.NewError(errorCode, msg)
}
