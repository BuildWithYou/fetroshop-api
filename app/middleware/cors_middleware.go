package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spf13/viper"
)

func CorsMiddleware(app *fiber.App, config *viper.Viper) {
	app.Use(cors.New(cors.Config{
		AllowOrigins: config.GetString("security.cors.allowOrigins"),
		AllowHeaders: config.GetString("security.cors.allowHeaders"),
		AllowMethods: config.GetString("security.cors.allowMethods"),
	}))
}
