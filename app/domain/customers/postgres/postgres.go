package postgres

import (
	"github.com/BuildWithYou/fetroshop-api/app/connection"
	"github.com/BuildWithYou/fetroshop-api/app/domain/customers"
	"gorm.io/gorm"
)

type PostgreSQL struct {
	Err error
	DB  *gorm.DB
}

func CustomerRepoProvider(db *connection.Connection) customers.CustomerRepo {
	return &PostgreSQL{
		Err: db.Err,
		DB:  db.DB,
	}
}
