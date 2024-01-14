package router

import (
	"github.com/BuildWithYou/fetroshop-api/app/helper/logger"
	"github.com/BuildWithYou/fetroshop-api/app/middleware"
	"github.com/BuildWithYou/fetroshop-api/app/modules/docs"
	"github.com/BuildWithYou/fetroshop-api/app/modules/web/controller"
	"github.com/gofiber/fiber/v2"
)

type WebRouter struct {
	Docs          *docs.Docs
	JwtMiddleware *middleware.JwtMiddleware
	DbMiddleware  *middleware.DbMiddleware
	Controller    *controller.Controller
	Logger        *logger.Logger
}

func (router *WebRouter) Init(app *fiber.App) {
	// Middlewares
	jwtMiddleware := router.JwtMiddleware.Authenticate
	dbMiddleware := router.DbMiddleware.Authenticate

	// root
	app.Get("/", router.redirectToDocs)
	app.Get("/welcome", router.welcome)

	// documentation
	app.Get("/documentation/*", router.Docs.SwaggerWeb())

	// api Group
	api := app.Group("/api", dbMiddleware)

	// Authentication
	authentication := api.Group("/auth")
	authentication.Post("/register", router.Controller.Auth.Register)
	authentication.Post("/login", router.Controller.Auth.Login)
	authentication.Post("/logout", jwtMiddleware, router.Controller.Auth.Logout)
	authentication.Post("/refresh", jwtMiddleware, router.Controller.Auth.Refresh)

	// Categories
	category := api.Group("/category")
	category.Get("/list", router.Controller.Category.List)
	category.Get("/find", router.Controller.Category.Find)

}

func WebRouterProvider(
	docs *docs.Docs,
	jwtMiddleware *middleware.JwtMiddleware,
	dbMiddleware *middleware.DbMiddleware,
	ctr *controller.Controller,
) Router {
	return &WebRouter{
		Docs:          docs,
		JwtMiddleware: jwtMiddleware,
		DbMiddleware:  dbMiddleware,
		Controller:    ctr,
	}
}

func (d *WebRouter) welcome(ctx *fiber.Ctx) error {
	return ctx.SendString("Welcome to fetroshop web api service!")
}

func (d *WebRouter) redirectToDocs(ctx *fiber.Ctx) error {
	return ctx.Redirect("/documentation/")
}
