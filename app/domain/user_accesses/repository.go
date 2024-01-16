package user_accesses

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserAccessRepo interface {
	Create(data *UserAccess) *gorm.DB
	Find(destination *UserAccess, condition fiber.Map) *gorm.DB
	UpdateOrCreate(data *UserAccess, condition fiber.Map) *gorm.DB
	Delete(condition *UserAccess) *gorm.DB
	Update(data *UserAccess, condition fiber.Map) *gorm.DB
}
