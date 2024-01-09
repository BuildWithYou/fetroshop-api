package postgres

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/customer_accesses"
	"gorm.io/gorm"
)

func (p *PostgreSQL) Create(data *customer_accesses.CustomerAccess) *gorm.DB {
	return p.DB.Create(data)
}
