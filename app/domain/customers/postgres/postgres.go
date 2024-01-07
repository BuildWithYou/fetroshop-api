package postgres

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/customers"
	"gorm.io/gorm"
)

type PostgreSQL struct {
	DB *gorm.DB
}

func CustomerRepoProvider(db *gorm.DB) customers.CustomerRepo {
	return &PostgreSQL{
		DB: db,
	}
}
