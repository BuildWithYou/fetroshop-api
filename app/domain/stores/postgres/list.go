package postgres

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/stores"
	"gorm.io/gorm"
)

func (p *PostgreSQL) List(destination *[]stores.Store, condition map[string]any, limit int, offset int, orderBy string) *gorm.DB {
	return p.DB.Where(condition).Limit(limit).Offset(offset).Order(orderBy).Find(destination)
}
