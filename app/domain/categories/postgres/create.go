package postgres

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/categories"
	"gorm.io/gorm"
)

func (p *PostgreSQL) Create(data *categories.Category) *gorm.DB {
	return p.DB.Create(data)
}
