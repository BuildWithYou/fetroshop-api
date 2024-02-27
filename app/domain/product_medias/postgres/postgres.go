package postgres

import (
	"github.com/BuildWithYou/fetroshop-api/app/connection"
	"github.com/BuildWithYou/fetroshop-api/app/domain/product_medias"
	"gorm.io/gorm"
)

type PostgreSQL struct {
	Err error
	DB  *gorm.DB
}

func RepoProvider(db *connection.Connection) product_medias.ProductMediaRepo {
	return &PostgreSQL{
		Err: db.Err,
		DB:  db.DB,
	}
}
