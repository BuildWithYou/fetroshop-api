package main

import (
	"time"

	"github.com/BuildWithYou/fetroshop-api/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
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

	app.Listen(":3000")
}
