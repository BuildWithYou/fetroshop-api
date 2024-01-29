package location

import (
	"github.com/BuildWithYou/fetroshop-api/app/connection"
	"github.com/BuildWithYou/fetroshop-api/app/domain/cities"
	"github.com/BuildWithYou/fetroshop-api/app/domain/districts"
	"github.com/BuildWithYou/fetroshop-api/app/domain/provinces"
	"github.com/BuildWithYou/fetroshop-api/app/domain/subdistricts"
	"github.com/BuildWithYou/fetroshop-api/app/helper/logger"
	"github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type LocationService interface {
	ListProvinces(ctx *fiber.Ctx) (*model.Response, error)
	ListCities(ctx *fiber.Ctx) (*model.Response, error)
	ListDistricts(ctx *fiber.Ctx) (*model.Response, error)
	ListSubdistricts(ctx *fiber.Ctx) (*model.Response, error)
}

type locationService struct {
	Err             error
	DB              *gorm.DB
	Config          *viper.Viper
	Validate        *validator.Validate
	ProvinceRepo    provinces.ProvinceRepo
	CityRepo        cities.CityRepo
	DistrictRepo    districts.DistrictRepo
	SubdistrictRepo subdistricts.SubdistrictRepo
	Logger          *logger.Logger
}

func ServiceProvider(
	conn *connection.Connection,
	config *viper.Viper,
	validate *validator.Validate,
	logger *logger.Logger,
	provRepo provinces.ProvinceRepo,
	cityRepo cities.CityRepo,
	districtRepo districts.DistrictRepo,
	subdistrictRepo subdistricts.SubdistrictRepo,
) LocationService {
	return &locationService{
		Err:             conn.Err,
		DB:              conn.DB,
		Config:          config,
		Validate:        validate,
		ProvinceRepo:    provRepo,
		CityRepo:        cityRepo,
		DistrictRepo:    districtRepo,
		SubdistrictRepo: subdistrictRepo,
		Logger:          logger,
	}
}
