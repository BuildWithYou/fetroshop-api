package auth

import (
	"github.com/BuildWithYou/fetroshop-api/app/connection"
	"github.com/BuildWithYou/fetroshop-api/app/domain/customer_accesses"
	"github.com/BuildWithYou/fetroshop-api/app/domain/customers"
	"github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

const CUSTOMER_TYPE = "customer"

type AuthService interface {
	Register(ctx *fiber.Ctx) (*model.Response, error)
	Login(ctx *fiber.Ctx) (*model.Response, error)
}

type AuthServiceImpl struct {
	DB                 *gorm.DB
	Config             *viper.Viper
	Validate           *validator.Validate
	CustomerRepo       customers.CustomerRepo
	CustomerAccessRepo customer_accesses.CustomerAccessRepo
}

func AuthServiceProvider(
	db *connection.Connection,
	config *viper.Viper,
	validate *validator.Validate,
	customerRepo customers.CustomerRepo,
	customerAccessRepo customer_accesses.CustomerAccessRepo,
) AuthService {
	return &AuthServiceImpl{
		DB:                 db.DB,
		Config:             config,
		Validate:           validate,
		CustomerRepo:       customerRepo,
		CustomerAccessRepo: customerAccessRepo,
	}
}
