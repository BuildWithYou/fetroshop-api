package main

import (
	"fmt"
	"time"

	"github.com/BuildWithYou/fetroshop-api/app/helper"
	"github.com/BuildWithYou/fetroshop-api/app/middleware"
	"github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/spf13/viper"
)

func main() {
	config := viper.New()
	config.SetConfigFile("config.json")
	err := config.ReadInConfig()
	helper.PanicIfError(err)

	app := fiber.New(fiber.Config{
		IdleTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 30,
		Prefork:      true,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			ctx.Status(fiber.StatusInternalServerError)
			return ctx.JSON(model.Response{
				Code:    fiber.ErrInternalServerError.Code,
				Status:  fiber.ErrInternalServerError.Message,
				Message: err.Error(),
			})

		},
	})

	app.Use(recover.New()) // Panic Handler

	// Routing
	ApiRoutes(app)

	// Middleware
	middleware.NotFoundMiddleware(app) // 404 Handler

	host := config.GetString("app.host")
	port := config.GetInt("app.port")
	err = app.Listen(fmt.Sprintf("%s:%d", host, port))
	helper.PanicIfError(err)
}
