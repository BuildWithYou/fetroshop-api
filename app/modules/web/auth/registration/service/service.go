package service

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/users"
	"github.com/gofiber/fiber/v2"
)

type RegistrationService interface {
	Register(ctx *fiber.Ctx) (err error)
}

type RegistrationServiceImpl struct {
	UserRepository users.UserRepository
}

func New(user users.UserRepository) RegistrationService {
	return &RegistrationServiceImpl{
		UserRepository: user,
	}
}
