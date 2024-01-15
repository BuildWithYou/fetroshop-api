package responsehelper

import (
	appModel "github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func Response200(msg string, data interface{}, meta interface{}) *appModel.Response {
	return &appModel.Response{
		Code:    fiber.StatusOK,
		Status:  utils.StatusMessage(fiber.StatusOK),
		Message: msg,
		Data:    data,
		Meta:    meta,
	}
}

func Response201(msg string, data interface{}, meta interface{}) *appModel.Response {
	return &appModel.Response{
		Code:    fiber.StatusCreated,
		Status:  utils.StatusMessage(fiber.StatusCreated),
		Message: msg,
		Data:    data,
		Meta:    meta,
	}
}

func Response400(msg string, data interface{}, meta interface{}) *appModel.Response {
	return &appModel.Response{
		Code:    fiber.StatusBadRequest,
		Status:  utils.StatusMessage(fiber.StatusBadRequest),
		Message: msg,
		Data:    data,
		Meta:    meta,
	}
}

func Response401(msg string, data interface{}, meta interface{}) *appModel.Response {
	return &appModel.Response{
		Code:    fiber.StatusUnauthorized,
		Status:  utils.StatusMessage(fiber.StatusUnauthorized),
		Message: msg,
		Data:    data,
		Meta:    meta,
	}
}

func Response500(msg string, data interface{}, meta interface{}) *appModel.Response {
	return &appModel.Response{
		Code:    fiber.StatusInternalServerError,
		Status:  utils.StatusMessage(fiber.StatusInternalServerError),
		Message: msg,
		Data:    data,
		Meta:    meta,
	}
}
