package category

import (
	"github.com/BuildWithYou/fetroshop-api/app/connection"
	"github.com/BuildWithYou/fetroshop-api/app/domain/categories"
	"github.com/BuildWithYou/fetroshop-api/app/helper/logger"
	"github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type CategoryService interface {
	Create(ctx *fiber.Ctx) (*model.Response, error)
	Update(ctx *fiber.Ctx) (*model.Response, error)
	Delete(ctx *fiber.Ctx) (*model.Response, error)
	List(ctx *fiber.Ctx) (*model.Response, error)
	Find(ctx *fiber.Ctx) (*model.Response, error)
}

type categoryService struct {
	Err          error
	DB           *gorm.DB
	Config       *viper.Viper
	Validate     *validator.Validate
	CategoryRepo categories.CategoryRepo
	Logger       *logger.Logger
}

func ServiceProvider(
	conn *connection.Connection,
	config *viper.Viper,
	validate *validator.Validate,
	logger *logger.Logger,
	categoryRepo categories.CategoryRepo,
) CategoryService {
	return &categoryService{
		Err:          conn.Err,
		DB:           conn.DB,
		Config:       config,
		Validate:     validate,
		CategoryRepo: categoryRepo,
		Logger:       logger,
	}
}
