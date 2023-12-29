package app

import (
	"github.com/BuildWithYou/fetroshop-api/app/modules"
	"github.com/BuildWithYou/fetroshop-api/app/modules/web/auth/registration/controller"
	"github.com/BuildWithYou/fetroshop-api/docs"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	WebRouter *WebRouter
	CmsRouter *CmsRouter
	Docs      *docs.Docs
}

type WebRouter struct {
	Registration controller.RegistrationController
}

type CmsRouter struct {
}

func (router *Router) Init(app *fiber.App) {
	// root
	app.Get("/", modules.Welcome)

	// swagger
	app.Get("/docs", router.Docs.Swagger)

	// registration
	app.Post("/web/api/register", router.WebRouter.Registration.Register)
	app.Get("/web/api/register", router.WebRouter.Registration.Register)
}
