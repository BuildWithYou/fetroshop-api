package router

import (
	"github.com/BuildWithYou/fetroshop-api/docs"
	"github.com/gofiber/fiber/v2"
)

type RouterImpl struct {
	WebRouter *WebRouter
	CmsRouter *CmsRouter
	Docs      *docs.Docs
}

type DocsRouter struct {
	Docs *docs.Docs
}

func DocsRouterProvider(d *docs.Docs) Router {
	return &DocsRouter{
		Docs: d,
	}
}

func (d *DocsRouter) Init(app *fiber.App) {
	// root
	app.Get("/", d.welcome)

	// docs
	app.Get("/web/*", d.Docs.SwaggerWeb())
	app.Get("/cms/*", d.Docs.SwaggerCms())
}

func (d *DocsRouter) welcome(ctx *fiber.Ctx) error {
	return ctx.SendString("Welcome to fetroshop-api docs!")
}
