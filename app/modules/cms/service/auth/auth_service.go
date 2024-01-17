package auth

import (
	"github.com/BuildWithYou/fetroshop-api/app/connection"
	"github.com/BuildWithYou/fetroshop-api/app/domain/user_accesses"
	"github.com/BuildWithYou/fetroshop-api/app/domain/users"
	"github.com/BuildWithYou/fetroshop-api/app/helper/constant"
	"github.com/BuildWithYou/fetroshop-api/app/helper/logger"
	appModel "github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/utils"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

const USER_TYPE = "user"

type AuthService interface {
	responseErrorGeneral(meta interface{}) *appModel.Response
	responseErrorValidation(meta interface{}) *appModel.Response
	Register(ctx *fiber.Ctx) (*appModel.Response, error)
	Login(ctx *fiber.Ctx) (*appModel.Response, error)
	Logout(ctx *fiber.Ctx) (*appModel.Response, error)
	Refresh(ctx *fiber.Ctx) (*appModel.Response, error)
}

type AuthServiceImpl struct {
	Err            error
	DB             *gorm.DB
	Config         *viper.Viper
	Validate       *validator.Validate
	UserRepo       users.UserRepo
	UserAccessRepo user_accesses.UserAccessRepo
	Logger         *logger.Logger
}

func ServiceProvider(
	conn *connection.Connection,
	config *viper.Viper,
	validate *validator.Validate,
	logger *logger.Logger,
	userRepo users.UserRepo,
	userAccessRepo user_accesses.UserAccessRepo,
) AuthService {
	return &AuthServiceImpl{
		Err:            conn.Err,
		DB:             conn.DB,
		Config:         config,
		Validate:       validate,
		UserRepo:       userRepo,
		UserAccessRepo: userAccessRepo,
		Logger:         logger,
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
