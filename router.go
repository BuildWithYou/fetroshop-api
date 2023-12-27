package main

import (
	"github.com/gofiber/fiber/v2"
)

func ApiRoutes(app *fiber.App) {
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Welcome to fetroshop-api!")
	})
}
