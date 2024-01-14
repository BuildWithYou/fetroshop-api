package router

import (
	"github.com/BuildWithYou/fetroshop-api/app/helper/logger"
	"github.com/BuildWithYou/fetroshop-api/app/middleware"
	"github.com/BuildWithYou/fetroshop-api/app/modules/cms/controller"
	"github.com/BuildWithYou/fetroshop-api/app/modules/docs"
	"github.com/gofiber/fiber/v2"
)

type CmsRouter struct {
	Docs          *docs.Docs
	JwtMiddleware *middleware.JwtMiddleware
	DbMiddleware  *middleware.DbMiddleware
	Controller    *controller.Controller
	Logger        *logger.Logger
}

func (router *CmsRouter) Init(app *fiber.App) {
	// Middlewares
	jwtMiddleware := router.JwtMiddleware.Authenticate
	dbMiddleware := router.DbMiddleware.Authenticate

	// root
	app.Get("/", router.redirectToDocs)
	app.Get("/welcome", router.welcome)

	// documentation
	app.Get("/documentation/*", router.Docs.SwaggerCms())

	// api Group
	api := app.Group("/api", dbMiddleware)

	// Authentication
	authentication := api.Group("/auth")
	authentication.Post("/register", router.Controller.Auth.Register)
	authentication.Post("/login", router.Controller.Auth.Login)
	authentication.Post("/logout", jwtMiddleware, router.Controller.Auth.Logout)
	authentication.Post("/refresh", jwtMiddleware, router.Controller.Auth.Refresh)

}

func CmsRouterProvider(
	docs *docs.Docs,
	jwtMiddleware *middleware.JwtMiddleware,
	dbMiddleware *middleware.DbMiddleware,
	ctr *controller.Controller,
) Router {
	return &CmsRouter{
		Docs:          docs,
		JwtMiddleware: jwtMiddleware,
		DbMiddleware:  dbMiddleware,
		Controller:    ctr,
	}
}

func (d *CmsRouter) welcome(ctx *fiber.Ctx) error {
	return ctx.SendString("Welcome to fetroshop cms api service!")
}

func (d *CmsRouter) redirectToDocs(ctx *fiber.Ctx) error {
	return ctx.Redirect("/documentation/")
}
