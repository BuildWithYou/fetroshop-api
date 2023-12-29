package modules

import (
	"github.com/gofiber/fiber/v2"
)

func Welcome(ctx *fiber.Ctx) error {
	return ctx.SendString("Welcome to fetroshop-api!")
}
