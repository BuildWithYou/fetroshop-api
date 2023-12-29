package main

import (
	"time"

	"github.com/BuildWithYou/fetroshop-api/app"
	"github.com/BuildWithYou/fetroshop-api/app/domain/users/postgres"
	"github.com/BuildWithYou/fetroshop-api/app/helper"
	"github.com/BuildWithYou/fetroshop-api/app/modules/web/controller"
	"github.com/BuildWithYou/fetroshop-api/app/modules/web/service/auth/registration"
	"github.com/BuildWithYou/fetroshop-api/docs"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func main() {
	// Config
	config := viper.New()
	config.SetConfigFile("config.json")
	err := config.ReadInConfig()
	helper.PanicIfError(err)

	// Fiber app initialization
	fiberApp := fiber.New(fiber.Config{
		IdleTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 30,
		Prefork:      true,
		ErrorHandler: helper.Error500,
	})

	// Postgres
	users := postgres.New()

	// Service
	registrationService := registration.New(users)

	// Controller
	registrationController := controller.New(registrationService)

	// Routing
	router := &app.Router{
		WebRouter: &app.WebRouter{
			Registration: registrationController,
		},
		CmsRouter: &app.CmsRouter{},
		Docs:      &docs.Docs{Config: config},
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
