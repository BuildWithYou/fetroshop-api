package auth

import (
	"github.com/BuildWithYou/fetroshop-api/app/connection"
	"github.com/BuildWithYou/fetroshop-api/app/domain/customer_accesses"
	"github.com/BuildWithYou/fetroshop-api/app/domain/customers"
	"github.com/BuildWithYou/fetroshop-api/app/domain/user_accesses"
	"github.com/BuildWithYou/fetroshop-api/app/domain/users"
	"github.com/BuildWithYou/fetroshop-api/app/helper/logger"
	appModel "github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

const USER_TYPE = "user"
const CUSTOMER_TYPE = "customer"

type AuthService interface {
	// Cms Modules
	CmsRegister(ctx *fiber.Ctx) (*appModel.Response, error)
	CmsLogin(ctx *fiber.Ctx) (*appModel.Response, error)
	CmsLogout(ctx *fiber.Ctx) (*appModel.Response, error)
	CmsRefresh(ctx *fiber.Ctx) (*appModel.Response, error)

	// Web Modules
	WebRegister(ctx *fiber.Ctx) (*appModel.Response, error)
	WebLogin(ctx *fiber.Ctx) (*appModel.Response, error)
	WebLogout(ctx *fiber.Ctx) (*appModel.Response, error)
	WebRefresh(ctx *fiber.Ctx) (*appModel.Response, error)
}

type authService struct {
	Err                error
	DB                 *gorm.DB
	Config             *viper.Viper
	Validate           *validator.Validate
	UserRepo           users.UserRepo
	UserAccessRepo     user_accesses.UserAccessRepo
	CustomerRepo       customers.CustomerRepo
	CustomerAccessRepo customer_accesses.CustomerAccessRepo
	Logger             *logger.Logger
}

func ServiceProvider(
	conn *connection.Connection,
	config *viper.Viper,
	validate *validator.Validate,
	userRepo users.UserRepo,
	userAccessRepo user_accesses.UserAccessRepo,
	customerRepo customers.CustomerRepo,
	customerAccessRepo customer_accesses.CustomerAccessRepo,
	logger *logger.Logger,
) AuthService {
	return &authService{
		Err:                conn.Err,
		DB:                 conn.DB,
		Config:             config,
		Validate:           validate,
		UserRepo:           userRepo,
		UserAccessRepo:     userAccessRepo,
		CustomerRepo:       customerRepo,
		CustomerAccessRepo: customerAccessRepo,
		Logger:             logger,
	}
}
