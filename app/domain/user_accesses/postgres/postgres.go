package postgres

import (
	"github.com/BuildWithYou/fetroshop-api/app/connection"
	"github.com/BuildWithYou/fetroshop-api/app/domain/user_accesses"
	"gorm.io/gorm"
)

type PostgreSQL struct {
	DB *gorm.DB
}

func UserAccessRepoProvider(db *connection.Connection) user_accesses.UserAccessRepo {
	return &PostgreSQL{
		DB: db.DB,
	}
}
