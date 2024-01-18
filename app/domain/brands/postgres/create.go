package postgres

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/brands"
	"gorm.io/gorm"
)

func (p *PostgreSQL) Create(data *brands.Brand) *gorm.DB {
	return p.DB.Create(data)
}
