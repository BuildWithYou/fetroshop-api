package router

import (
	"github.com/BuildWithYou/fetroshop-api/app/modules/web/controller"
	"github.com/gofiber/fiber/v2"
)

type WebRouter struct {
	Registration controller.RegistrationController
}

func (router *WebRouter) Init(app *fiber.App) {
	// root
	app.Get("/", router.welcome)

	// registration
	app.Post("/api/auth/register", router.Registration.Register)
}

func WebRouterProvider(ctr controller.RegistrationController) Router {
	return &WebRouter{
		Registration: ctr,
	}
}

func (d *WebRouter) welcome(ctx *fiber.Ctx) error {
	return ctx.SendString("Welcome to fetroshop web api service!")
}
