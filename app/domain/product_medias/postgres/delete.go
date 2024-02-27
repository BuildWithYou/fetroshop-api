package postgres

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/product_medias"
	"gorm.io/gorm"
)

func (p *PostgreSQL) Delete(condition map[string]any) *gorm.DB {
	return p.DB.Where(condition).Delete(&product_medias.ProductMedia{})
}
