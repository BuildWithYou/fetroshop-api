package postgres

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/cities"
	"github.com/BuildWithYou/fetroshop-api/app/helper/gormhelper"
	"gorm.io/gorm"
)

func (p *PostgreSQL) Find(destination *cities.City, condition map[string]any) *gorm.DB {
	query := gormhelper.ConditionMapping(p.DB, condition)
	if query.Error != nil {
		return query
	}

	return query.First(destination)
}
