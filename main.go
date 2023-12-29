package main

import (
	"time"

	"github.com/BuildWithYou/fetroshop-api/app"
	"github.com/BuildWithYou/fetroshop-api/app/domain/users/postgres"
	"github.com/BuildWithYou/fetroshop-api/app/helper"
	"github.com/BuildWithYou/fetroshop-api/app/modules/web/auth/registration/controller"
	"github.com/BuildWithYou/fetroshop-api/app/modules/web/auth/registration/service"
	"github.com/BuildWithYou/fetroshop-api/docs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/spf13/viper"
)

// @Version 1.0.0
// @Title Fetroshop API
// @Description Fetroshop API is a robust and efficient backend solution designed to power the online store app named Fetroshop. Developed using the Go programming language, this API serves as the backbone for managing the Content Management System (CMS) and handling various store-related functionalities.
// @ContactName Eko Teguh Wicaksono
// @ContactEmail ekoteguhwicaksono@gmail.com
// @Security AuthorizationHeader read write
// @SecurityScheme AuthorizationHeader apiKey header X-MyCustomHeader
func main() {
	// Config
	config := viper.New()
	config.SetConfigFile("config.json")
	err := config.ReadInConfig()
	helper.PanicIfError(err)

	// View engine
	viewEngine := html.New("./app/views", ".html")

	// Fiber app initialization
	fiberApp := fiber.New(fiber.Config{
		IdleTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 30,
		Views:        viewEngine,
		Prefork:      true,
		ErrorHandler: helper.Error500,
	})

	// Postgres
	users := postgres.New()

	// Service
	registrationService := service.New(users)

	// Controller
	registrationController := controller.New(registrationService)

	// Routing
	router := &app.Router{
		WebRouter: &app.WebRouter{
			Registration: registrationController,
		},
		CmsRouter: &app.CmsRouter{},
		Docs: &docs.Docs{
			Config: config},
	}

	// Initialize Fetroshop App
	fetroshopApp := app.App{
		Config:   config,
		FiberApp: fiberApp,
		Router:   router,
	}
	err = fetroshopApp.Start()
	helper.PanicIfError(err)
}
