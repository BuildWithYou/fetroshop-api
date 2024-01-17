package postgres

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/categories"
	"gorm.io/gorm"
)

func (p *PostgreSQL) Update(data *categories.Category, condition map[string]any) *gorm.DB {
	return p.DB.Where(condition).Updates(data)
}
