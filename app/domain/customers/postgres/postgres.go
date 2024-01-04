package postgres

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/customers"
	"gorm.io/gorm"
)

type PostgreSQL struct {
	DB *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) customers.CustomerRepository {
	return &PostgreSQL{
		DB: db,
	}
}
