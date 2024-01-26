package postgres

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/stores"
	"gorm.io/gorm"
)

func (p *PostgreSQL) Create(data *stores.Store) *gorm.DB {
	return p.DB.Create(data)
}
