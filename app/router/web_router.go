package router

import (
	"github.com/BuildWithYou/fetroshop-api/app/middleware"
	"github.com/BuildWithYou/fetroshop-api/app/modules/docs"
	"github.com/BuildWithYou/fetroshop-api/app/modules/web/controller"
	"github.com/gofiber/fiber/v2"
)

type WebRouter struct {
	Docs          *docs.Docs
	JwtMiddleware *middleware.JwtMiddleware
	Controller    *controller.Controller
}

func (router *WebRouter) Init(app *fiber.App) {
	// Middlewares
	jwtMiddleware := router.JwtMiddleware.Authenticate

	// root
	app.Get("/", router.welcome)

	// documentation
	app.Get("/documentation/*", router.Docs.SwaggerWeb())

	// Authentication
	authentication := app.Group("/api/auth")
	authentication.Post("/register", router.Controller.Auth.Register)
	authentication.Post("/login", router.Controller.Auth.Login)
	authentication.Post("/logout", jwtMiddleware, router.Controller.Auth.Logout)
	authentication.Post("/refresh", jwtMiddleware, router.Controller.Auth.Refresh)

}

func WebRouterProvider(
	docs *docs.Docs,
	ctr *controller.Controller,
	jwtMiddleware *middleware.JwtMiddleware,
) Router {
	return &WebRouter{
		Docs:          docs,
		Controller:    ctr,
		JwtMiddleware: jwtMiddleware,
	}
}

func (d *WebRouter) welcome(ctx *fiber.Ctx) error {
	return ctx.SendString("Welcome to fetroshop web api service!")
}
