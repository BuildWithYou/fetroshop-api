package categories

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type CategoryRepo interface {
	Create(data *Category) *gorm.DB
	Find(destination *Category, condition map[string]any) *gorm.DB
	List(destination *[]Category, condition fiber.Map, limit int, offset int, orderBy string) *gorm.DB
}
