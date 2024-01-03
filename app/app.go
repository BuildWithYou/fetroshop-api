package app

import (
	"fmt"
	"time"

	"github.com/BuildWithYou/fetroshop-api/app/helper"
	"github.com/BuildWithYou/fetroshop-api/app/middleware"
	"github.com/BuildWithYou/fetroshop-api/app/router"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/spf13/viper"
)

type App struct {
	Config     *viper.Viper
	FiberApp   *fiber.App
	Router     router.Router
	Validation *validator.Validate
}

func (app *App) Start() error {

	if app.Config.GetBool("fiber.recovery") {
		app.FiberApp.Use(recover.New(recover.Config{
			EnableStackTrace: app.Config.GetBool("fiber.enableStackTrace"),
		})) // Panic Handler
	}

	app.Router.Init(app.FiberApp)

	// Swagger static files
	app.FiberApp.Static("/swagger", "docs")

	// Middleware
	middleware.NotFoundMiddleware(app.FiberApp) // 404 Handler

	host := app.Config.GetString("app.web.host")
	port := app.Config.GetInt("app.web.port")
	return app.FiberApp.Listen(fmt.Sprintf("%s:%d", host, port))
}

type ServerConfig struct {
	Config *viper.Viper
	Host   string
	Port   int
	Router router.Router
}

func GetConfig() *viper.Viper {
	// Config
	config := viper.New()
	config.SetConfigFile("config.yaml")
	err := config.ReadInConfig()
	helper.PanicIfError(err)
	return config
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

	// Middleware
	middleware.NotFoundMiddleware(fiberApp) // 404 Handler

	host := serverConfig.Host
	port := serverConfig.Port
	return fiberApp.Listen(fmt.Sprintf("%s:%d", host, port))
}

func DocsServerConfigProvider(webRouter router.Router) *ServerConfig {
	config := GetConfig()
	return &ServerConfig{
		Config: config,
		Host:   config.GetString("app.docs.host"),
		Port:   config.GetInt("app.docs.port"),
		Router: webRouter,
	}
}
