package postgres

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/customer_accesses"
	"gorm.io/gorm"
)

func (p *PostgreSQL) Update(data *customer_accesses.CustomerAccess, condition map[string]any) *gorm.DB {
	return p.DB.Where(condition).Updates(data)
}
