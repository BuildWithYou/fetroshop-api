package postgres

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/user_accesses"
	"gorm.io/gorm"
)

func (p *PostgreSQL) Create(data *user_accesses.UserAccess) *gorm.DB {
	return p.DB.Create(data)
}
