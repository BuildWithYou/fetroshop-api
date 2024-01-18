package postgres

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/brands"
	"gorm.io/gorm"
)

func (p *PostgreSQL) Delete(condition map[string]any) *gorm.DB {
	return p.DB.Where(condition).Delete(&brands.Brand{})
}
