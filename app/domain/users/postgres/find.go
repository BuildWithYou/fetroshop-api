package postgres

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/users"
	"gorm.io/gorm"
)

func (p *PostgreSQL) Find(destination *users.User, condition *users.User) *gorm.DB {
	return p.DB.Where(condition).First(destination)
}
