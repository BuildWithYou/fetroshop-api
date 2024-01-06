package postgres

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/customer_accesses"
	"gorm.io/gorm"
)

func (p *PostgreSQL) Find(destination *customer_accesses.CustomerAccess, condition *customer_accesses.CustomerAccess) *gorm.DB {
	return p.DB.Where(condition).First(destination)
}
