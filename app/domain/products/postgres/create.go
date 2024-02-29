package postgres

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/products"
	"gorm.io/gorm"
)

func (p *PostgreSQL) Create(data *products.Product) *gorm.DB {
	return p.DB.Create(data)
}
