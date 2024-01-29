package postgres

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/cities"
	"github.com/BuildWithYou/fetroshop-api/app/helper/gormhelper"
	"gorm.io/gorm"
)

func (p *PostgreSQL) Count(destination *int64, condition map[string]any) *gorm.DB {
	query := p.DB.Model(&cities.City{})
	query = gormhelper.ConditionMapping(query, condition)
	if query.Error != nil {
		return query
	}

	return query.Count(destination)
}
