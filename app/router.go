package app

import (
	"github.com/BuildWithYou/fetroshop-api/app/modules"
	"github.com/BuildWithYou/fetroshop-api/app/modules/web/controller"
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

	// docs
	app.Get("/docs/web/*", router.Docs.SwaggerWeb())
	app.Get("/docs/cms/*", router.Docs.SwaggerCms())

	// registration
	app.Post("/api/web/register", router.WebRouter.Registration.Register)
}
