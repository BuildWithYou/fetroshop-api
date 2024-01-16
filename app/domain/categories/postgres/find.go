package postgres

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/categories"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func (p *PostgreSQL) Find(destination *categories.Category, condition fiber.Map) *gorm.DB {
	return p.DB.Preload("Parent").Where(condition).First(destination)
}
