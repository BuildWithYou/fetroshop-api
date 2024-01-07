package postgres

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/customer_accesses"
	"gorm.io/gorm"
)

type PostgreSQL struct {
	DB *gorm.DB
}

func CustomerAccessRepoProvider(db *gorm.DB) customer_accesses.CustomerAccessRepo {
	return &PostgreSQL{
		DB: db,
	}
}
