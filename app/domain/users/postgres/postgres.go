package postgres

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/users"
	"gorm.io/gorm"
)

type PostgreSQL struct {
	DB *gorm.DB
}

func UserRepositoryProvider(db *gorm.DB) users.UserRepository {
	return &PostgreSQL{
		DB: db,
	}
}
