package postgres

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/products"
	"gorm.io/gorm"
)

func (p *PostgreSQL) Delete(condition map[string]any) *gorm.DB {
	return p.DB.Where(condition).Delete(&products.Product{})
}
