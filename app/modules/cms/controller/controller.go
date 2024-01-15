package controller

import (
	appModel "github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	Auth AuthController
}

func CmsControllerProvider(
	authController AuthController,
) *Controller {
	return &Controller{
		Auth: authController,
	}
}

func execute(ctx *fiber.Ctx, handler func(ctx *fiber.Ctx) (*appModel.Response, error)) (err error) {
	response, err := handler(ctx)
	if err != nil {
		return err
	}
	return ctx.JSON(response)
}
