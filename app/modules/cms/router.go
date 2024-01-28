package cms

import (
	"github.com/BuildWithYou/fetroshop-api/app"
	"github.com/BuildWithYou/fetroshop-api/app/helper/logger"
	"github.com/BuildWithYou/fetroshop-api/app/middleware"
	"github.com/BuildWithYou/fetroshop-api/app/modules/cms/controller"
	"github.com/BuildWithYou/fetroshop-api/app/modules/docs"
	"github.com/gofiber/fiber/v2"
)

type CmsRouter struct {
	Docs             *docs.Docs
	JwtMiddleware    *middleware.JwtMiddleware
	DbMiddleware     *middleware.DbMiddleware
	LoggerMiddleware *middleware.LoggerMiddleware
	Controller       *controller.Controller
}

func (router *CmsRouter) Init(app *fiber.App) {
	// Middlewares
	jwtMiddleware := router.JwtMiddleware.Authenticate
	dbMiddleware := router.DbMiddleware.Authenticate
	loggerMiddleware := router.LoggerMiddleware.CmsLoggerResetOutput
	contentTypeMiddleware := middleware.ContentTypeMiddleware

	// root
	app.Get("/", router.redirectToDocs)
	app.Get("/welcome", router.welcome)

	// documentation
	app.Get("/documentation/*", router.Docs.SwaggerCms())

	// api Group
	api := app.Group("/api", dbMiddleware, loggerMiddleware)

	// Authentication
	authentication := api.Group("/auth")
	authentication.Post("/register", contentTypeMiddleware, router.Controller.Auth.Register)
	authentication.Post("/login", contentTypeMiddleware, router.Controller.Auth.Login)
	authentication.Post("/logout", jwtMiddleware, router.Controller.Auth.Logout)
	authentication.Post("/refresh", jwtMiddleware, router.Controller.Auth.Refresh)

	// Category
	category := api.Group("/category")
	category.Post("/create", contentTypeMiddleware, jwtMiddleware, router.Controller.Category.Create)
	category.Put("/:code", contentTypeMiddleware, jwtMiddleware, router.Controller.Category.Update)
	category.Delete("/:code", contentTypeMiddleware, jwtMiddleware, router.Controller.Category.Delete)
	category.Get("/list", router.Controller.Category.List)
	category.Get("/detail", router.Controller.Category.Find)

	// Brand
	brand := api.Group("/brand")
	brand.Post("/create", contentTypeMiddleware, jwtMiddleware, router.Controller.Brand.Create)
	brand.Put("/:code", contentTypeMiddleware, jwtMiddleware, router.Controller.Brand.Update)
	brand.Delete("/:code", contentTypeMiddleware, jwtMiddleware, router.Controller.Brand.Delete)
	brand.Get("/list", router.Controller.Brand.List)
	brand.Get("/list-by-prefix", router.Controller.Brand.ListByPrefix)
	brand.Get("/detail", router.Controller.Brand.Find)

	// Store
	store := api.Group("/store")
	store.Post("/create", contentTypeMiddleware, jwtMiddleware, router.Controller.Store.Create)
	store.Put("/:code", contentTypeMiddleware, jwtMiddleware, router.Controller.Store.Update)
	store.Delete("/:code", contentTypeMiddleware, jwtMiddleware, router.Controller.Store.Delete)
	store.Get("/list", router.Controller.Store.List)
	store.Get("/detail", router.Controller.Store.Find)

	// Location
	location := api.Group("/location")
	location.Get("/province/list", router.Controller.Location.ListProvinces)
	location.Get("/province/detail", router.Controller.Location.FindProvince)
	location.Get("/city/list", router.Controller.Location.ListCities)
	location.Get("/city/detail", router.Controller.Location.FindCity)
	location.Get("/district/list", router.Controller.Location.ListDistricts)
	location.Get("/district/detail", router.Controller.Location.FindDistrict)
	location.Get("/subdistrict/list", router.Controller.Location.ListSubdistricts)
	location.Get("/subdistrict/detail", router.Controller.Location.FindSubdistrict)
}

func RouterProvider(
	docs *docs.Docs,
	jwtMiddleware *middleware.JwtMiddleware,
	dbMiddleware *middleware.DbMiddleware,
	ctr *controller.Controller,
	logger *logger.Logger,
) app.Router {
	loggerMiddleware := middleware.LoggerMiddlewareProvider(logger)
	return &CmsRouter{
		Docs:             docs,
		JwtMiddleware:    jwtMiddleware,
		DbMiddleware:     dbMiddleware,
		LoggerMiddleware: loggerMiddleware,
		Controller:       ctr,
	}
}

func (d *CmsRouter) welcome(ctx *fiber.Ctx) error {
	return ctx.SendString("Welcome to fetroshop cms api service!")
}

func (d *CmsRouter) redirectToDocs(ctx *fiber.Ctx) error {
	return ctx.Redirect("/documentation/")
}
