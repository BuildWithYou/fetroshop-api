package postgres

import (
	"github.com/BuildWithYou/fetroshop-api/app/connection"
	"github.com/BuildWithYou/fetroshop-api/app/domain/subdistricts"
	"gorm.io/gorm"
)

type PostgreSQL struct {
	Err error
	DB  *gorm.DB
}

func RepoProvider(db *connection.Connection) subdistricts.SubdistrictRepo {
	return &PostgreSQL{
		Err: db.Err,
		DB:  db.DB,
	}
}
