package service

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/users"
	"github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/BuildWithYou/fetroshop-api/app/modules/web/auth/registration"
)

type RegistrationService interface {
	Register(*registration.RegistrationRequest) (*model.GeneralResponse, error)
}

type RegistrationServiceImpl struct {
	UserRepository users.UserRepository
}

func New(user users.UserRepository) RegistrationService {
	return &RegistrationServiceImpl{
		UserRepository: user,
	}
}
