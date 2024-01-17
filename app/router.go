package app

import "github.com/gofiber/fiber/v2"

type Router interface {
	Init(app *fiber.App)
}
