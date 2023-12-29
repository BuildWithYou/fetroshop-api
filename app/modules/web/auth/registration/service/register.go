package service

import (
	"github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/BuildWithYou/fetroshop-api/app/modules/web/auth/registration"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/utils"
)

func (rg *RegistrationServiceImpl) Register(*registration.RegistrationRequest) (*model.GeneralResponse, error) {
	// TODO - Implement Register
	return &model.GeneralResponse{
		Code:    fiber.StatusCreated,
		Status:  utils.StatusMessage(fiber.StatusCreated),
		Message: "Not Implemented",
	}, nil
}
