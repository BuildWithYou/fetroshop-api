package controller

import (
	"github.com/BuildWithYou/fetroshop-api/app/modules/auth/registration/service"
	"github.com/gofiber/fiber/v2"
)

type RegistrationController interface {
	Register(ctx *fiber.Ctx) (err error)
}

type RegistrationControllerImpl struct{}

func New(registrationService service.RegistrationService) RegistrationController {
	return &RegistrationControllerImpl{}
}
