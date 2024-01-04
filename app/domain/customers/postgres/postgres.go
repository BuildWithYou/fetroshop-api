package postgres

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/customers"
)

type PostgreSQL struct{}

func NewCustomerRepository() customers.CustomerRepository {
	return &PostgreSQL{}
}
