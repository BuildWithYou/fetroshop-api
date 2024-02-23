package postgres

import (
	"strings"

	"github.com/BuildWithYou/fetroshop-api/app/domain/products"
	"gorm.io/gorm"
)

func (p *PostgreSQL) Count(destination *int64, keyword string) *gorm.DB {
	query := p.DB.Model(&products.Product{}).Where("UPPER(name) LIKE ?", "%"+strings.ToUpper(keyword)+"%")
	if query.Error != nil {
		return query
	}

	return query.Count(destination)
}
