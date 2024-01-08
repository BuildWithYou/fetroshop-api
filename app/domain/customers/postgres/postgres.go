package postgres

import (
	"github.com/BuildWithYou/fetroshop-api/app/connection"
	"github.com/BuildWithYou/fetroshop-api/app/domain/customers"
	"gorm.io/gorm"
)

type PostgreSQL struct {
	DB *gorm.DB
}

func CustomerRepoProvider(db *connection.Connection) customers.CustomerRepo {
	return &PostgreSQL{
		DB: db.DB,
	}
}
