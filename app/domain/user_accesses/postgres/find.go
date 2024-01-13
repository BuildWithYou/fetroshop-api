package postgres

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/user_accesses"
	"gorm.io/gorm"
)

func (p *PostgreSQL) Find(destination *user_accesses.UserAccess, condition map[string]any) *gorm.DB {
	return p.DB.Where(condition).First(destination)
}
