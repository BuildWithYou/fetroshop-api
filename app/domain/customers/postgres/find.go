package postgres

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/customers"
	"gorm.io/gorm"
)

func (p *PostgreSQL) Find(destination *customers.Customer, condition map[string]any) *gorm.DB {
	return p.DB.Where(condition).First(destination)
}
