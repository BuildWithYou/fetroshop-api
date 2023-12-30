package router

import (
	"github.com/BuildWithYou/fetroshop-api/docs"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	WebRouter *WebRouter
	CmsRouter *CmsRouter
	Docs      *docs.Docs
}

func (router *Router) Init(app *fiber.App) {
	// root
	app.Get("/", welcome)

	// docs
	app.Get("/docs/web/*", router.Docs.SwaggerWeb())
	app.Get("/docs/cms/*", router.Docs.SwaggerCms())
}

func welcome(ctx *fiber.Ctx) error {
	return ctx.SendString("Welcome to fetroshop-api!")
}
