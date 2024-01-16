package customer_accesses

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type CustomerAccessRepo interface {
	Create(data *CustomerAccess) *gorm.DB
	UpdateOrCreate(data *CustomerAccess, condition fiber.Map) *gorm.DB
	Update(data *CustomerAccess, condition fiber.Map) *gorm.DB
	Find(destination *CustomerAccess, condition fiber.Map) *gorm.DB
	Delete(condition *CustomerAccess) *gorm.DB
}
