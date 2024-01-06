package router

import (
	"github.com/BuildWithYou/fetroshop-api/app/middleware"
	"github.com/BuildWithYou/fetroshop-api/app/modules/cms/controller"
	"github.com/BuildWithYou/fetroshop-api/app/modules/docs"
	"github.com/gofiber/fiber/v2"
)

type CmsRouter struct {
	Docs          *docs.Docs
	JwtMiddleware *middleware.JwtMiddleware
	Controller    *controller.Controller
}

func (router *CmsRouter) Init(app *fiber.App) {
	// root
	app.Get("/", router.welcome)

	// documentation
	app.Get("/documentation/*", router.Docs.SwaggerCms())

	// Authentication
	authentication := app.Group("/api/auth")
	authentication.Post("/register", router.Controller.Auth.Register)
	authentication.Post("/login", router.Controller.Auth.Login)
}

func CmsRouterProvider(docs *docs.Docs, jwtMiddleware *middleware.JwtMiddleware, ctr *controller.Controller) Router {
	return &CmsRouter{
		Docs:          docs,
		JwtMiddleware: jwtMiddleware,
		Controller:    ctr,
	}
}

func (d *CmsRouter) welcome(ctx *fiber.Ctx) error {
	return ctx.SendString("Welcome to fetroshop cms api service!")
}
