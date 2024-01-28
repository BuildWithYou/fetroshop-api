package postgres

import (
	"strings"

	"github.com/BuildWithYou/fetroshop-api/app/domain/stores"
	"gorm.io/gorm"
)

func (p *PostgreSQL) Count(destination *int64, keyword string) *gorm.DB {
	query := p.DB.Model(&stores.Store{}).Where("UPPER(name) LIKE ?", "%"+strings.ToUpper(keyword)+"%").
		Or("UPPER(code) LIKE ?", "%"+strings.ToUpper(keyword)+"%")
	if query.Error != nil {
		return query
	}

	return query.Count(destination)
}
