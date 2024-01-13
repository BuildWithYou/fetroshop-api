package postgres

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/categories"
	"gorm.io/gorm"
)

func (p *PostgreSQL) Find(destination *categories.Category, condition map[string]any) *gorm.DB {
	return p.DB.Preload("Parent").Where(condition).First(destination)
}
