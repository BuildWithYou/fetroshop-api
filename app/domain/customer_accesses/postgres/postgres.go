package postgres

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/customer_accesses"
	"gorm.io/gorm"
)

type PostgreSQL struct {
	DB *gorm.DB
}

func CustomerAccessRepositoryProvider(db *gorm.DB) customer_accesses.CustomerAccessRepository {
	return &PostgreSQL{
		DB: db,
	}
}
