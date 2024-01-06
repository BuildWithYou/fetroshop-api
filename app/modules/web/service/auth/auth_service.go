package auth

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/customers"
	"github.com/BuildWithYou/fetroshop-api/app/model"
	webModel "github.com/BuildWithYou/fetroshop-api/app/modules/web/model"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type AuthService interface {
	Register(request *webModel.RegistrationRequest) (*model.Response, error)
	Login(request *webModel.LoginRequest) (*model.Response, error)
}

type AuthServiceImpl struct {
	DB                 *gorm.DB
	Config             *viper.Viper
	CustomerRepository customers.CustomerRepository
}

func AuthServiceProvider(
	db *gorm.DB,
	config *viper.Viper,
	customerRepository customers.CustomerRepository) AuthService {
	return &AuthServiceImpl{
		DB:                 db,
		Config:             config,
		CustomerRepository: customerRepository,
	}
}
