package app

import (
	"github.com/BuildWithYou/fetroshop-api/app/modules/auth/registration/controller"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	Registration controller.RegistrationController
}

func (router *Router) ApiRoutes(app *fiber.App) {
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Welcome to fetroshop-api!")
	})

	app.Post("/api/v1/register", router.Registration.Register)
	app.Get("/api/v1/register", router.Registration.Register)
}
