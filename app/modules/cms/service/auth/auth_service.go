package auth

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/user_accesses"
	"github.com/BuildWithYou/fetroshop-api/app/domain/users"
	"github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

const USER_TYPE = "user"

type AuthService interface {
	Register(ctx *fiber.Ctx) (*model.Response, error)
	Login(ctx *fiber.Ctx) (*model.Response, error)
}

type AuthServiceImpl struct {
	DB             *gorm.DB
	Config         *viper.Viper
	Validate       *validator.Validate
	UserRepo       users.UserRepo
	UserAccessRepo user_accesses.UserAccessRepo
}

func AuthServiceProvider(
	db *gorm.DB,
	config *viper.Viper,
	validate *validator.Validate,
	userRepo users.UserRepo,
	userAccessRepo user_accesses.UserAccessRepo,
) AuthService {
	return &AuthServiceImpl{
		DB:             db,
		Config:         config,
		Validate:       validate,
		UserRepo:       userRepo,
		UserAccessRepo: userAccessRepo,
	}
}
