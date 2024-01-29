package postgres

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/stores"
	"github.com/BuildWithYou/fetroshop-api/app/helper/gormhelper"
	"gorm.io/gorm"
)

func (p *PostgreSQL) FindWithLocation(destination *stores.Store, condition map[string]any) *gorm.DB {
	query := p.DB.Preload("Province").Preload("City").Preload("District").Preload("Subdistrict")
	query = gormhelper.ConditionMapping(query, condition)
	if query.Error != nil {
		return query
	}

	return query.First(destination)
}
