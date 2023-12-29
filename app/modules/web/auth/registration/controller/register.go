package controller

import (
	"github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/gofiber/fiber/v2"
)

func (r *RegistrationControllerImpl) Register(ctx *fiber.Ctx) (err error) {
	// TODO - Implement Register
	return ctx.JSON(model.Response{
		Code:    fiber.ErrInternalServerError.Code,
		Status:  fiber.ErrInternalServerError.Message,
		Message: "Not Implemented",
	})
}
