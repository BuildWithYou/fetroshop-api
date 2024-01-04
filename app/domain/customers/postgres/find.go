package postgres

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/customers"
)

func (p *PostgreSQL) Find(condition *customers.Customer) *customers.Customer {
	destination := new(customers.Customer)
	p.DB.Where(condition).First(destination)
	return destination
}
