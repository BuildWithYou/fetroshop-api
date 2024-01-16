package postgres

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/customer_accesses"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func (p *PostgreSQL) Find(destination *customer_accesses.CustomerAccess, condition fiber.Map) *gorm.DB {
	return p.DB.Where(condition).First(destination)
}
