package postgres

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/product_medias"
	"gorm.io/gorm"
)

func (p *PostgreSQL) Create(data *product_medias.ProductMedia) *gorm.DB {
	return p.DB.Create(data)
}
