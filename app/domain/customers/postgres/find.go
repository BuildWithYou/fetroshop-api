package postgres

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/customers"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func (p *PostgreSQL) Find(destination *customers.Customer, condition fiber.Map) *gorm.DB {
	return p.DB.Where(condition).First(destination)
}
