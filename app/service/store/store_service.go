package store

import (
	"github.com/BuildWithYou/fetroshop-api/app/connection"
	"github.com/BuildWithYou/fetroshop-api/app/domain/stores"
	"github.com/BuildWithYou/fetroshop-api/app/helper/logger"
	appModel "github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type StoreService interface {
	Create(ctx *fiber.Ctx) (*appModel.Response, error)
	Update(ctx *fiber.Ctx) (*appModel.Response, error)
	Delete(ctx *fiber.Ctx) (*appModel.Response, error)
	List(ctx *fiber.Ctx) (*appModel.Response, error)
	Find(ctx *fiber.Ctx) (*appModel.Response, error)
}

type storeService struct {
	Err       error
	DB        *gorm.DB
	Config    *viper.Viper
	Validate  *validator.Validate
	StoreRepo stores.StoreRepo
	Logger    *logger.Logger
}

func ServiceProvider(
	conn *connection.Connection,
	config *viper.Viper,
	validate *validator.Validate,
	logger *logger.Logger,
	storeRepo stores.StoreRepo,
) StoreService {
	return &storeService{
		Err:       conn.Err,
		DB:        conn.DB,
		Config:    config,
		Validate:  validate,
		StoreRepo: storeRepo,
		Logger:    logger,
	}
}
