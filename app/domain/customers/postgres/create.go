package postgres

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/customers"
	"gorm.io/gorm"
)

func (p *PostgreSQL) Create(cst *customers.Customer) *gorm.DB {
	return p.DB.Create(cst)
}
