package app

import (
	"errors"
	"fmt"
	"time"

	"github.com/BuildWithYou/fetroshop-api/app/helper/logger"
	"github.com/BuildWithYou/fetroshop-api/app/middleware"
	appModel "github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/BuildWithYou/fetroshop-api/app/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/utils"
	"github.com/spf13/viper"
)

type Fetroshop struct {
	FiberApp *fiber.App
	Host     string
	Port     int
	Err      error
	Logger   *logger.Logger
}

type ServerConfig struct {
	Config *viper.Viper
	Host   string
	Port   int
	Router router.Router
	Static map[string]string
	Logger *logger.Logger
}

// CreateFiber initializes a Fiber app with the given server configuration and returns a Fetroshop instance.
//
// The function takes a pointer to a ServerConfig struct as its parameter.
// The ServerConfig struct contains various settings for the Fiber app, such as idle timeout, write timeout, read timeout, prefork, and error handler.
//
// The function returns a pointer to a Fetroshop struct, which contains the initialized Fiber app, host, and port.
func CreateFiber(serverConfig *ServerConfig) *Fetroshop {
	// Fiber app initialization
	fiberApp := fiber.New(fiber.Config{
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

			if code == fiber.StatusInternalServerError {
				logger := logger.NewFrameworkLogger()
				logger.Error(fmt.Sprint("Error : ", err.Error()))
			}

			return ctx.Status(code).JSON(appModel.Response{
				Code:    code,
				Status:  status,
				Message: err.Error(),
			})
		},
	})

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
	return &Fetroshop{
		FiberApp: fiberApp,
		Host:     host,
		Port:     port,
	}
}
