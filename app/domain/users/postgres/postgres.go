package postgres

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/users"
	"gorm.io/gorm"
)

type PostgreSQL struct {
	DB *gorm.DB
}

func UserRepoProvider(db *gorm.DB) users.UserRepo {
	return &PostgreSQL{
		DB: db,
	}
}
