package customers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type CustomerRepo interface {
	Create(data *Customer) *gorm.DB
	Find(destination *Customer, condition fiber.Map) *gorm.DB
}
