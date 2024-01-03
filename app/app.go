package app

import (
	"fmt"
	"time"

	"github.com/BuildWithYou/fetroshop-api/app/helper"
	"github.com/BuildWithYou/fetroshop-api/app/middleware"
	"github.com/BuildWithYou/fetroshop-api/app/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/spf13/viper"
)

type ServerConfig struct {
	Config *viper.Viper
	Host   string
	Port   int
	Router router.Router
	Static map[string]string
}

func CreateFiber(serverConfig *ServerConfig) (fiberApp *fiber.App) {
	// Fiber app initialization
	return fiber.New(fiber.Config{
		IdleTimeout:  time.Second * time.Duration(serverConfig.Config.GetInt("fiber.idleTimeout")),
		WriteTimeout: time.Second * time.Duration(serverConfig.Config.GetInt("fiber.writeTimeout")),
		ReadTimeout:  time.Second * time.Duration(serverConfig.Config.GetInt("fiber.readTimeout")),
		Prefork:      serverConfig.Config.GetBool("fiber.prefork"),
		ErrorHandler: helper.ErrorCustom,
	})
}

func StartFiber(
	fiberApp *fiber.App,
	serverConfig *ServerConfig) error {

	if serverConfig.Config.GetBool("fiber.recovery") {
		fiberApp.Use(recover.New(recover.Config{
			EnableStackTrace: serverConfig.Config.GetBool("fiber.enableStackTrace"),
		})) // Panic Handler
	}

	serverConfig.Router.Init(fiberApp)

	// Static files
	if serverConfig.Static != nil {
		for key, value := range serverConfig.Static {
			fiberApp.Static(key, value)
		}
	}

	// Middleware
	middleware.NotFoundMiddleware(fiberApp) // 404 Handler

	host := serverConfig.Host
	port := serverConfig.Port
	return fiberApp.Listen(fmt.Sprintf("%s:%d", host, port))
}
