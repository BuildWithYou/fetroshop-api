package errorhelper

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func PanicIfError(err error) {
	if err != nil {
		fmt.Println("Error : ", err.Error()) // #marked: logging
		panic(err)
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
