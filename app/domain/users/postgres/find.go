package postgres

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/users"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func (p *PostgreSQL) Find(destination *users.User, condition fiber.Map) *gorm.DB {
	return p.DB.Where(condition).First(destination)
}
