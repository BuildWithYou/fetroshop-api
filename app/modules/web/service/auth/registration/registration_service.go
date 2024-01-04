package registration

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/customers"
	"github.com/BuildWithYou/fetroshop-api/app/model"
	webModel "github.com/BuildWithYou/fetroshop-api/app/modules/web/model"
	"gorm.io/gorm"
)

type RegistrationService interface {
	Register(request *webModel.RegistrationRequest) (*model.Response, error)
}

type RegistrationServiceImpl struct {
	DB                 *gorm.DB
	CustomerRepository customers.CustomerRepository
}

func NewRegistrationService(db *gorm.DB, customerRepository customers.CustomerRepository) RegistrationService {
	return &RegistrationServiceImpl{
		DB:                 db,
		CustomerRepository: customerRepository,
	}
}
