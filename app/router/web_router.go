package router

import (
	"github.com/BuildWithYou/fetroshop-api/app/middleware"
	"github.com/BuildWithYou/fetroshop-api/app/modules/docs"
	"github.com/BuildWithYou/fetroshop-api/app/modules/web/controller"
	"github.com/gofiber/fiber/v2"
)

type WebRouter struct {
	Docs           *docs.Docs
	JwtMiddleware  *middleware.JwtMiddleware
	Authentication controller.AuthController
}

func (router *WebRouter) Init(app *fiber.App) {
	// root
	app.Get("/", router.welcome)

	// documentation
	app.Get("/documentation/*", router.Docs.SwaggerWeb())

	// Authentication
	app.Post("/api/auth/register", router.Authentication.Register)
	app.Post("/api/auth/login", router.Authentication.Login)

}

func WebRouterProvider(
	docs *docs.Docs,
	ctr controller.AuthController,
	jwtMiddleware *middleware.JwtMiddleware,
) Router {
	return &WebRouter{
		Docs:           docs,
		Authentication: ctr,
		JwtMiddleware:  jwtMiddleware,
	}
}

func (d *WebRouter) welcome(ctx *fiber.Ctx) error {
	return ctx.SendString("Welcome to fetroshop web api service!")
}
