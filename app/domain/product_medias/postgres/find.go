package postgres

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/product_medias"
	"github.com/BuildWithYou/fetroshop-api/app/helper/gormhelper"
	"gorm.io/gorm"
)

func (p *PostgreSQL) Find(destination *product_medias.ProductMedia, condition map[string]any) *gorm.DB {
	query := gormhelper.ConditionMapping(p.DB, condition)
	if query.Error != nil {
		return query
	}

	return query.First(destination)
}
