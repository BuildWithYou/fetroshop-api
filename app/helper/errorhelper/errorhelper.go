package errorhelper

import (
	"encoding/json"
	"fmt"

	"github.com/BuildWithYou/fetroshop-api/app/helper/logger"
	"github.com/gofiber/fiber/v2"
)

var fwLogger = logger.NewFrameworkLogger()

func generateErrorJson(jsonable interface{}) string {
	// make json string
	errorJson, err := json.Marshal(jsonable)
	if err != nil {
		return "Failed to parse error message to json"
	}
	errorMesage := string(errorJson)
	return errorMesage
}

func Error400(jsonable interface{}) error {
	errorMesage := generateErrorJson(jsonable)
	return fiber.NewError(fiber.StatusBadRequest, errorMesage)
}

func Error401(jsonable interface{}) error {
	errorMesage := generateErrorJson(jsonable)
	return fiber.NewError(fiber.StatusUnauthorized, errorMesage)
}

func Error500(jsonable interface{}) error {
	errorMesage := generateErrorJson(jsonable)
	fwLogger.Error(fmt.Sprint("Error : ", errorMesage))
	return fiber.NewError(fiber.StatusInternalServerError, errorMesage)
}

func ErrorCustom(errorCode int, jsonable interface{}) error {
	errorMesage := generateErrorJson(jsonable)
	return fiber.NewError(errorCode, errorMesage)
}
