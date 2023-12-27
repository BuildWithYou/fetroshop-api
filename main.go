package main

import (
	"fmt"
	"time"

	"github.com/BuildWithYou/fetroshop-api/helper"
	"github.com/BuildWithYou/fetroshop-api/model/api"
	"github.com/BuildWithYou/fetroshop-api/routes"
	"github.com/gofiber/fiber/v2"
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
	})

	app.Use(func(c *fiber.Ctx) error {
		err := c.Next()
		return err
	})

	routes.ApiRoutes(app)

	// 404 Handler
	app.Use(func(ctx *fiber.Ctx) error {
		return ctx.JSON(api.Response{
			Code:   fiber.ErrNotFound.Code,
			Status: fiber.ErrNotFound.Message,
		})
	})

	host := config.GetString("app.host")
	port := config.GetInt("app.port")
	err = app.Listen(fmt.Sprintf("%s:%d", host, port))
	helper.PanicIfError(err)
}
