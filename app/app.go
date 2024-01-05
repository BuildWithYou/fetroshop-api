package app

import (
	"errors"
	"fmt"
	"time"

	"github.com/BuildWithYou/fetroshop-api/app/middleware"
	"github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/BuildWithYou/fetroshop-api/app/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/utils"
	"github.com/spf13/viper"
)

type ServerConfig struct {
	Config *viper.Viper
	Host   string
	Port   int
	Router router.Router
	Static map[string]string
}

// CreateFiber initializes a Fiber app with the given server configuration.
//
// Parameters:
// - serverConfig: a pointer to a ServerConfig object containing the server configuration.
//
// Returns:
// - fiberApp: a pointer to a fiber.App object representing the initialized Fiber app.
func CreateFiber(serverConfig *ServerConfig) (fiberApp *fiber.App) {
	// Fiber app initialization
	return fiber.New(fiber.Config{
		IdleTimeout:  time.Second * time.Duration(serverConfig.Config.GetInt("fiber.idleTimeout")),
		WriteTimeout: time.Second * time.Duration(serverConfig.Config.GetInt("fiber.writeTimeout")),
		ReadTimeout:  time.Second * time.Duration(serverConfig.Config.GetInt("fiber.readTimeout")),
		Prefork:      serverConfig.Config.GetBool("fiber.prefork"),
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			// Status code defaults to 500
			code := fiber.StatusInternalServerError
			status := fiber.ErrInternalServerError.Message

			// Retrieve the custom status code if it's a *fiber.Error
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
				status = utils.StatusMessage(e.Code)
			}

			return ctx.Status(code).JSON(model.Response{
				Code:    code,
				Status:  status,
				Message: err.Error(),
			})
		},
	})
}

// StartFiber initializes and starts a Fiber application.
//
// fiberApp: A pointer to the Fiber application instance.
// serverConfig: A pointer to the server configuration.
// Returns an error if there is an issue starting the Fiber application.
func StartFiber(
	fiberApp *fiber.App,
	serverConfig *ServerConfig) error {

	middleware.CorsMiddleware(fiberApp, serverConfig.Config)

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
