package postgres

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/user_accesses"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func (p *PostgreSQL) Find(destination *user_accesses.UserAccess, condition fiber.Map) *gorm.DB {
	return p.DB.Where(condition).First(destination)
}
