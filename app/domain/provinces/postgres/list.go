package postgres

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/provinces"
	"github.com/BuildWithYou/fetroshop-api/app/helper/gormhelper"
	"gorm.io/gorm"
)

func (p *PostgreSQL) List(destination *[]provinces.Province, condition map[string]any, limit int, offset int, orderBy string) *gorm.DB {
	query := gormhelper.ConditionMapping(p.DB, condition)
	if query.Error != nil {
		return query
	}

	if limit > 0 {
		query = query.Limit(limit).Offset(offset)
	}

	return query.Order(orderBy).Find(destination)
}
