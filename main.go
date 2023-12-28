package main

import (
	"fmt"
	"time"

	"github.com/BuildWithYou/fetroshop-api/app"
	"github.com/BuildWithYou/fetroshop-api/app/domain/users/postgres"
	"github.com/BuildWithYou/fetroshop-api/app/helper"
	"github.com/BuildWithYou/fetroshop-api/app/middleware"
	"github.com/BuildWithYou/fetroshop-api/app/modules/auth/registration/controller"
	"github.com/BuildWithYou/fetroshop-api/app/modules/auth/registration/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/spf13/viper"
)

func main() {
	config := viper.New()
	config.SetConfigFile("config.json")
	err := config.ReadInConfig()
	helper.PanicIfError(err)

	fiberApp := fiber.New(fiber.Config{
		IdleTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 30,
		Prefork:      true,
		ErrorHandler: helper.Error500,
	})

	fiberApp.Use(recover.New()) // Panic Handler

	// Postgres
	users := postgres.New()

	// Service
	registrationService := service.New(users)

	// Controller
	registrationController := controller.New(registrationService)

	// Routing
	router := app.Router{
		Registration: registrationController,
	}
	router.ApiRoutes(fiberApp)

	// Middleware
	middleware.NotFoundMiddleware(fiberApp) // 404 Handler

	host := config.GetString("app.host")
	port := config.GetInt("app.port")
	err = fiberApp.Listen(fmt.Sprintf("%s:%d", host, port))
	helper.PanicIfError(err)
}
