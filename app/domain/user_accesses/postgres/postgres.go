package postgres

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/user_accesses"
	"gorm.io/gorm"
)

type PostgreSQL struct {
	DB *gorm.DB
}

func UserAccessRepositoryProvider(db *gorm.DB) user_accesses.UserAccessRepository {
	return &PostgreSQL{
		DB: db,
	}
}
