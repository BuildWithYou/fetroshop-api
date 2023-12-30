package main

import (
	"time"

	"github.com/BuildWithYou/fetroshop-api/app"
	"github.com/BuildWithYou/fetroshop-api/app/domain/users/postgres"
	"github.com/BuildWithYou/fetroshop-api/app/helper"
	"github.com/BuildWithYou/fetroshop-api/app/modules/web/controller"
	"github.com/BuildWithYou/fetroshop-api/app/modules/web/service/auth/registration"
	"github.com/BuildWithYou/fetroshop-api/app/router"
	"github.com/BuildWithYou/fetroshop-api/docs"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func main() {
	// Config
	config := viper.New()
	config.SetConfigFile("config.yaml")
	err := config.ReadInConfig()
	helper.PanicIfError(err)

	// Validation
	validate := validator.New()

	// Fiber app initialization
	fiberApp := fiber.New(fiber.Config{
		IdleTimeout:  time.Second * time.Duration(config.GetInt("fiber.idleTimeout")),
		WriteTimeout: time.Second * time.Duration(config.GetInt("fiber.writeTimeout")),
		ReadTimeout:  time.Second * time.Duration(config.GetInt("fiber.readTimeout")),
		Prefork:      config.GetBool("fiber.prefork"),
		ErrorHandler: helper.ErrorCustom,
	})

	// Postgres
	users := postgres.New()

	// Service
	registrationService := registration.New(&registration.RegistrationServiceTransport{
		Validation:     validate,
		UserRepository: users,
	})

	// Controller
	registrationController := controller.New(&controller.RegistrationControllerTransport{
		Validate:            validate,
		RegistrationService: registrationService,
	})

	// Routing
	router := &router.Router{
		WebRouter: &router.WebRouter{
			Registration: registrationController,
		},
		CmsRouter: &router.CmsRouter{},
		Docs:      &docs.Docs{Config: config},
	}

	// Initialize Fetroshop App
	fetroshopApp := app.App{
		Config:     config,
		FiberApp:   fiberApp,
		Router:     router,
		Validation: validate,
	}
	err = fetroshopApp.Start()
	helper.PanicIfError(err)
}
