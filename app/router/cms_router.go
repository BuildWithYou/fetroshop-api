package router

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type CmsRouter struct {
	Validation *validator.Validate
}

func (router *CmsRouter) Init(app *fiber.App) {
	// cms router
}
