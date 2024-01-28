package controller

import (
	"github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	Auth     AuthController
	Category CategoryController
	Brand    BrandController
	Location LocationController
}

func WebControllerProvider(
	authController AuthController,
	categoryController CategoryController,
	brandController BrandController,
	locationController LocationController,
) *Controller {
	return &Controller{
		Auth:     authController,
		Category: categoryController,
		Brand:    brandController,
		Location: locationController,
	}
}

func execute(ctx *fiber.Ctx, handler func(ctx *fiber.Ctx) (*model.Response, error)) (err error) {
	response, err := handler(ctx)
	if err != nil {
		return err
	}
	return ctx.Status(response.Code).JSON(response)
}
