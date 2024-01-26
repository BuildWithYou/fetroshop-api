package postgres

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/cities"
	"github.com/BuildWithYou/fetroshop-api/app/helper/gormhelper"
	"gorm.io/gorm"
)

func (p *PostgreSQL) List(destination *[]cities.City, condition map[string]any, limit int, offset int, orderBy string) *gorm.DB {
	query := gormhelper.ConditionMapping(p.DB, condition)
	if query.Error != nil {
		return query
	}

	return query.Limit(limit).Offset(offset).Order(orderBy).Find(destination)
}