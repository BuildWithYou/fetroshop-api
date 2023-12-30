package registration

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/users"
	"github.com/BuildWithYou/fetroshop-api/app/model"
	webModel "github.com/BuildWithYou/fetroshop-api/app/modules/web/model"
	"github.com/go-playground/validator/v10"
)

type RegistrationService interface {
	Register(request *webModel.RegistrationRequest) (*model.Response, error)
}

type RegistrationServiceTransport struct {
	Validation     *validator.Validate
	UserRepository users.UserRepository
}

type RegistrationServiceV1 struct {
	Validation     *validator.Validate
	UserRepository users.UserRepository
}

func New(o *RegistrationServiceTransport) RegistrationService {
	return &RegistrationServiceV1{
		UserRepository: o.UserRepository,
	}
}
