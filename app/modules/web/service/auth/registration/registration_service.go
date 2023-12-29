package registration

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/users"
	"github.com/BuildWithYou/fetroshop-api/app/model"
	webModel "github.com/BuildWithYou/fetroshop-api/app/modules/web/model"
)

type RegistrationService interface {
	Register(*webModel.RegistrationRequest) (*model.Response, error)
}

type RegistrationServiceImpl struct {
	UserRepository users.UserRepository
}

func New(user users.UserRepository) RegistrationService {
	return &RegistrationServiceImpl{
		UserRepository: user,
	}
}
