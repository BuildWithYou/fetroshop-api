package postgres

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/categories"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func (p *PostgreSQL) List(destination *[]categories.Category, condition fiber.Map, limit int, offset int, orderBy string) *gorm.DB {
	return p.DB.Preload("Parent").Where(condition).Limit(limit).Offset(offset).Order(orderBy).Find(destination)
}
