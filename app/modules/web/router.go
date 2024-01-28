package web

import (
	"github.com/BuildWithYou/fetroshop-api/app"
	"github.com/BuildWithYou/fetroshop-api/app/helper/logger"
	"github.com/BuildWithYou/fetroshop-api/app/middleware"
	"github.com/BuildWithYou/fetroshop-api/app/modules/docs"
	"github.com/BuildWithYou/fetroshop-api/app/modules/web/controller"
	"github.com/gofiber/fiber/v2"
)

type WebRouter struct {
	Docs             *docs.Docs
	JwtMiddleware    *middleware.JwtMiddleware
	DbMiddleware     *middleware.DbMiddleware
	LoggerMiddleware *middleware.LoggerMiddleware
	Controller       *controller.Controller
}

func (router *WebRouter) Init(app *fiber.App) {
	// Middlewares
	jwtMiddleware := router.JwtMiddleware.Authenticate
	dbMiddleware := router.DbMiddleware.Authenticate
	loggerMiddleware := router.LoggerMiddleware.WebLoggerResetOutput
	contentTypeMiddleware := middleware.ContentTypeMiddleware

	// root
	app.Get("/", router.redirectToDocs)
	app.Get("/welcome", router.welcome)

	// documentation
	app.Get("/documentation/*", router.Docs.SwaggerWeb())

	// api Group
	api := app.Group("/api", dbMiddleware, loggerMiddleware)

	// Authentication
	authentication := api.Group("/auth")
	authentication.Post("/register", contentTypeMiddleware, router.Controller.Auth.Register)
	authentication.Post("/login", contentTypeMiddleware, router.Controller.Auth.Login)
	authentication.Post("/logout", jwtMiddleware, router.Controller.Auth.Logout)
	authentication.Post("/refresh", jwtMiddleware, router.Controller.Auth.Refresh)

	// Categories
	category := api.Group("/category")
	category.Get("/list", router.Controller.Category.List)
	category.Get("/detail", router.Controller.Category.Find)

	// Brand
	brand := api.Group("/brand")
	brand.Get("/list", router.Controller.Brand.List)
	brand.Get("/list-by-prefix", router.Controller.Brand.ListByPrefix)
	brand.Get("/detail", router.Controller.Brand.Find)

	// Location
	location := api.Group("/location")
	location.Get("/province/list", router.Controller.Location.ListProvinces)
	location.Get("/city/list", router.Controller.Location.ListCities)
	location.Get("/district/list", router.Controller.Location.ListDistricts)
	location.Get("/subdistrict/list", router.Controller.Location.ListSubdistricts)
}

func RouterProvider(
	docs *docs.Docs,
	jwtMiddleware *middleware.JwtMiddleware,
	dbMiddleware *middleware.DbMiddleware,
	ctr *controller.Controller,
	logger *logger.Logger,
) app.Router {
	loggerMiddleware := middleware.LoggerMiddlewareProvider(logger)
	return &WebRouter{
		Docs:             docs,
		JwtMiddleware:    jwtMiddleware,
		DbMiddleware:     dbMiddleware,
		LoggerMiddleware: loggerMiddleware,
		Controller:       ctr,
	}
}

func (d *WebRouter) welcome(ctx *fiber.Ctx) error {
	return ctx.SendString("Welcome to fetroshop web api service!")
}

func (d *WebRouter) redirectToDocs(ctx *fiber.Ctx) error {
	return ctx.Redirect("/documentation/")
}
