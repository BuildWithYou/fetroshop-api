package middleware

import (
	"github.com/BuildWithYou/fetroshop-api/app/model/api"
	"github.com/gofiber/fiber/v2"
)

func NotFoundMiddleware(app *fiber.App) {
	app.Use(func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusNotFound).JSON(api.Response{
			Code:    fiber.ErrNotFound.Code,
			Status:  fiber.ErrNotFound.Message,
			Message: "Sorry, page not found!",
		})
	})
}
