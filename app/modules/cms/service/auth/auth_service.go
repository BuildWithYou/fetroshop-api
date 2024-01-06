package auth

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/users"
	"github.com/BuildWithYou/fetroshop-api/app/model"
	webModel "github.com/BuildWithYou/fetroshop-api/app/modules/cms/model"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type AuthService interface {
	Register(request *webModel.RegistrationRequest) (*model.Response, error)
	Login(request *webModel.LoginRequest) (*model.Response, error)
}

type AuthServiceImpl struct {
	DB             *gorm.DB
	Config         *viper.Viper
	UserRepository users.UserRepository
}

func AuthServiceProvider(
	db *gorm.DB,
	config *viper.Viper,
	userRepository users.UserRepository) AuthService {
	return &AuthServiceImpl{
		DB:             db,
		Config:         config,
		UserRepository: userRepository,
	}
}
