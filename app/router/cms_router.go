package router

import (
	"github.com/gofiber/fiber/v2"
)

type CmsRouter struct {
}

func (router *CmsRouter) Init(app *fiber.App) {
	// root
	app.Get("/", router.welcome)
}

func CmsRouterProvider() Router {
	return &CmsRouter{}
}

func (d *CmsRouter) welcome(ctx *fiber.Ctx) error {
	return ctx.SendString("Welcome to fetroshop cms api service!")
}
