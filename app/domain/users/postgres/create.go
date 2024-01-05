package postgres

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/users"
	"gorm.io/gorm"
)

func (p *PostgreSQL) Create(user *users.User) *gorm.DB {
	return p.DB.Create(user)
}
