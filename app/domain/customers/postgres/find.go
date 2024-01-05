package postgres

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/customers"
	"gorm.io/gorm"
)

func (p *PostgreSQL) Find(destination *customers.Customer, condition *customers.Customer) *gorm.DB {
	return p.DB.Where(condition).First(destination)
}
