package location

import (
	"github.com/BuildWithYou/fetroshop-api/app/connection"
	"github.com/BuildWithYou/fetroshop-api/app/domain/stores"
	"github.com/BuildWithYou/fetroshop-api/app/helper/logger"
	"github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type LocationService interface {
	ListProvinces(ctx *fiber.Ctx) (*model.Response, error)
	FindProvince(ctx *fiber.Ctx) (*model.Response, error)
	ListCities(ctx *fiber.Ctx) (*model.Response, error)
	FindCity(ctx *fiber.Ctx) (*model.Response, error)
	ListDistricts(ctx *fiber.Ctx) (*model.Response, error)
	FindDistrict(ctx *fiber.Ctx) (*model.Response, error)
	ListSubdistricts(ctx *fiber.Ctx) (*model.Response, error)
	FindSubdistrict(ctx *fiber.Ctx) (*model.Response, error)
}

type locationService struct {
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
) LocationService {
	return &locationService{
		Err:       conn.Err,
		DB:        conn.DB,
		Config:    config,
		Validate:  validate,
		StoreRepo: storeRepo,
		Logger:    logger,
	}
}
