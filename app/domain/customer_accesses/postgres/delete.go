package postgres

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/customer_accesses"
	"gorm.io/gorm"
)

func (p *PostgreSQL) Delete(condition *customer_accesses.CustomerAccess) *gorm.DB {
	return p.DB.Delete(condition)
}
