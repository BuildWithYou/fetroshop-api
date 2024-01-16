package postgres

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/customer_accesses"
	"gorm.io/gorm"
)

func (p *PostgreSQL) Find(destination *customer_accesses.CustomerAccess, condition map[string]any) *gorm.DB {
	return p.DB.Where(condition).First(destination)
}
