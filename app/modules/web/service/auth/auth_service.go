package auth

import (
	"github.com/BuildWithYou/fetroshop-api/app/connection"
	"github.com/BuildWithYou/fetroshop-api/app/domain/customer_accesses"
	"github.com/BuildWithYou/fetroshop-api/app/domain/customers"
	"github.com/BuildWithYou/fetroshop-api/app/helper/constant"
	"github.com/BuildWithYou/fetroshop-api/app/helper/logger"
	appModel "github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/utils"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

const CUSTOMER_TYPE = "customer"

type AuthService interface {
	responseErrorGeneral(meta interface{}) *appModel.Response
	responseErrorValidation(meta interface{}) *appModel.Response
	Register(ctx *fiber.Ctx) (*appModel.Response, error)
	Login(ctx *fiber.Ctx) (*appModel.Response, error)
	Logout(ctx *fiber.Ctx) (*appModel.Response, error)
	Refresh(ctx *fiber.Ctx) (*appModel.Response, error)
}

type AuthServiceImpl struct {
	Err                error
	DB                 *gorm.DB
	Config             *viper.Viper
	Validate           *validator.Validate
	CustomerRepo       customers.CustomerRepo
	CustomerAccessRepo customer_accesses.CustomerAccessRepo
	Logger             *logger.Logger
}

func ServiceProvider(
	conn *connection.Connection,
	config *viper.Viper,
	validate *validator.Validate,
	logger *logger.Logger,
	customerRepo customers.CustomerRepo,
	customerAccessRepo customer_accesses.CustomerAccessRepo,
) AuthService {
	return &AuthServiceImpl{
		Err:                conn.Err,
		DB:                 conn.DB,
		Config:             config,
		Validate:           validate,
		CustomerRepo:       customerRepo,
		CustomerAccessRepo: customerAccessRepo,
		Logger:             logger,
	}
}

func (svc *AuthServiceImpl) responseErrorGeneral(meta interface{}) *appModel.Response {
	svc.Logger.Error(meta)
	return &appModel.Response{
		Code:    fiber.StatusInternalServerError,
		Status:  utils.StatusMessage(fiber.StatusInternalServerError),
		Message: constant.ERROR_GENERAL,
		Meta:    meta,
	}
}

func (svc *AuthServiceImpl) responseErrorValidation(meta interface{}) *appModel.Response {
	return &appModel.Response{
		Code:    fiber.StatusBadRequest,
		Status:  utils.StatusMessage(fiber.StatusBadRequest),
		Message: constant.ERROR_VALIDATION,
		Meta:    meta,
	}
}
