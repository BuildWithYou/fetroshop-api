package users

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserRepo interface {
	Create(data *User) *gorm.DB
	Find(destination *User, condition fiber.Map) *gorm.DB
}
