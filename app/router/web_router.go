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
	WebController *controller.WebController
}

func (router *WebRouter) Init(app *fiber.App) {
	// root
	app.Get("/", router.welcome)

	// documentation
	app.Get("/documentation/*", router.Docs.SwaggerWeb())

	// Authentication
	authentication := app.Group("/api/auth")
	authentication.Post("/register", router.WebController.Auth.Register)
	authentication.Post("/login", router.WebController.Auth.Login)

}

func WebRouterProvider(
	docs *docs.Docs,
	ctr *controller.WebController,
	jwtMiddleware *middleware.JwtMiddleware,
) Router {
	return &WebRouter{
		Docs:          docs,
		WebController: ctr,
		JwtMiddleware: jwtMiddleware,
	}
}

func (d *WebRouter) welcome(ctx *fiber.Ctx) error {
	return ctx.SendString("Welcome to fetroshop web api service!")
}
