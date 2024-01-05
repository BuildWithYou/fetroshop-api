package router

import (
	"github.com/BuildWithYou/fetroshop-api/app/middleware"
	"github.com/BuildWithYou/fetroshop-api/app/modules/docs"
	"github.com/gofiber/fiber/v2"
)

type CmsRouter struct {
	Docs          *docs.Docs
	JwtMiddleware *middleware.JwtMiddleware
}

func (router *CmsRouter) Init(app *fiber.App) {
	// root
	app.Get("/", router.welcome)

	// documentation
	app.Get("/documentation/*", router.Docs.SwaggerCms())
}

func CmsRouterProvider(docs *docs.Docs, jwtMiddleware *middleware.JwtMiddleware) Router {
	return &CmsRouter{
		Docs:          docs,
		JwtMiddleware: jwtMiddleware,
	}
}

func (d *CmsRouter) welcome(ctx *fiber.Ctx) error {
	return ctx.SendString("Welcome to fetroshop cms api service!")
}
