package controller

import (
	appModel "github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	Auth     AuthController
	Category CategoryController
	Brand    brandController
}

func CmsControllerProvider(
	authController AuthController,
	categoryController CategoryController,
) *Controller {
	return &Controller{
		Auth:     authController,
		Category: categoryController,
	}
}

func execute(ctx *fiber.Ctx, handler func(ctx *fiber.Ctx) (*appModel.Response, error)) (err error) {
	response, err := handler(ctx)
	if err != nil {
		return err
	}
	return ctx.Status(response.Code).JSON(response)
}
