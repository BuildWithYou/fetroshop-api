package postgres

import (
	"strings"

	"github.com/BuildWithYou/fetroshop-api/app/domain/stores"
	"gorm.io/gorm"
)

func (p *PostgreSQL) List(destination *[]stores.Store, keyword string, limit int, offset int, orderBy string) *gorm.DB {
	query := p.DB.Where("UPPER(name) LIKE ?", "%"+strings.ToUpper(keyword)+"%").
		Or("UPPER(code) LIKE ?", "%"+strings.ToUpper(keyword)+"%")
	if query.Error != nil {
		return query
	}

	if limit > 0 {
		query = query.Limit(limit).Offset(offset)
	}

	return query.Order(orderBy).Find(destination)
}
