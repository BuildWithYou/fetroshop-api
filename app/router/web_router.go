package router

import (
	"github.com/BuildWithYou/fetroshop-api/app/modules/web/controller"
	"github.com/gofiber/fiber/v2"
)

type WebRouter struct {
	Registration controller.RegistrationController
}

func (router *WebRouter) Init(app *fiber.App) {
	// registration
	app.Post("/api/web/register", router.Registration.Register)
}