package controller

import (
	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
	appModel "github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	Auth     AuthController
	Category CategoryController
}

func WebControllerProvider(
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
	if validatorhelper.IsNotNil(err) {
		return err
	}
	return ctx.JSON(response)
}
