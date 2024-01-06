package auth

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/customers"
	"github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type AuthService interface {
	Register(ctx *fiber.Ctx) (*model.Response, error)
	Login(ctx *fiber.Ctx) (*model.Response, error)
}

type AuthServiceImpl struct {
	DB                 *gorm.DB
	Config             *viper.Viper
	Validate           *validator.Validate
	CustomerRepository customers.CustomerRepository
}

func AuthServiceProvider(
	db *gorm.DB,
	config *viper.Viper,
	validate *validator.Validate,
	customerRepository customers.CustomerRepository) AuthService {
	return &AuthServiceImpl{
		DB:                 db,
		Config:             config,
		Validate:           validate,
		CustomerRepository: customerRepository,
	}
}
