package app

import (
	"fmt"

	"github.com/BuildWithYou/fetroshop-api/app/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/spf13/viper"
)

type App struct {
	Config   *viper.Viper
	FiberApp *fiber.App
	Router   *Router
}

func (app *App) Start() error {

	app.FiberApp.Use(recover.New()) // Panic Handler

	app.Router.Init(app.FiberApp)

	// Swagger static files
	app.FiberApp.Static("/swagger", "docs")

	// Middleware
	middleware.NotFoundMiddleware(app.FiberApp) // 404 Handler

	host := app.Config.GetString("app.host")
	port := app.Config.GetInt("app.port")
	return app.FiberApp.Listen(fmt.Sprintf("%s:%d", host, port))
}
