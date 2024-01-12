package controller

import (
	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	Auth AuthController
}

func WebControllerProvider(
	authController AuthController,
) *Controller {
	return &Controller{
		Auth: authController,
	}
}

func execute(ctx *fiber.Ctx, handler func(ctx *fiber.Ctx) (*model.Response, error)) (err error) {
	response, err := handler(ctx)
	if validatorhelper.IsNotNil(err) {
		return err
	}
	return ctx.JSON(response)
}
