package store

import (
	"github.com/BuildWithYou/fetroshop-api/app/connection"
	"github.com/BuildWithYou/fetroshop-api/app/domain/cities"
	"github.com/BuildWithYou/fetroshop-api/app/domain/districts"
	"github.com/BuildWithYou/fetroshop-api/app/domain/provinces"
	"github.com/BuildWithYou/fetroshop-api/app/domain/stores"
	"github.com/BuildWithYou/fetroshop-api/app/domain/subdistricts"
	"github.com/BuildWithYou/fetroshop-api/app/helper/logger"
	"github.com/BuildWithYou/fetroshop-api/app/helper/miniohelper"
	"github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type StoreService interface {
	Create(ctx *fiber.Ctx) (*model.Response, error)
	Update(ctx *fiber.Ctx) (*model.Response, error)
	Delete(ctx *fiber.Ctx) (*model.Response, error)
	List(ctx *fiber.Ctx) (*model.Response, error)
	Find(ctx *fiber.Ctx) (*model.Response, error)
}

type storeService struct {
	Err             error
	DB              *gorm.DB
	Config          *viper.Viper
	Validate        *validator.Validate
	StoreRepo       stores.StoreRepo
	ProvinceRepo    provinces.ProvinceRepo
	CityRepo        cities.CityRepo
	DistrictRepo    districts.DistrictRepo
	SubdistrictRepo subdistricts.SubdistrictRepo
	Logger          *logger.Logger
	Minio           *miniohelper.Minio
}

func ServiceProvider(
	conn *connection.Connection,
	config *viper.Viper,
	validate *validator.Validate,
	storeRepo stores.StoreRepo,
	provRepo provinces.ProvinceRepo,
	cityRepo cities.CityRepo,
	districtRepo districts.DistrictRepo,
	subdistrictRepo subdistricts.SubdistrictRepo,
	logger *logger.Logger,
	myMinio *miniohelper.Minio,
) StoreService {
	return &storeService{
		Err:             conn.Err,
		DB:              conn.DB,
		Config:          config,
		Validate:        validate,
		StoreRepo:       storeRepo,
		ProvinceRepo:    provRepo,
		CityRepo:        cityRepo,
		DistrictRepo:    districtRepo,
		SubdistrictRepo: subdistrictRepo,
		Logger:          logger,
		Minio:           myMinio,
	}
}
