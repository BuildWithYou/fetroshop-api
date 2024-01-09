package postgres

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/user_accesses"
	"gorm.io/gorm"
)

func (p *PostgreSQL) Delete(condition *user_accesses.UserAccess) *gorm.DB {
	return p.DB.Delete(&user_accesses.UserAccess{}, condition)
}
