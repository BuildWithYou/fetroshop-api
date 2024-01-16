package postgres

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/user_accesses"
	"github.com/BuildWithYou/fetroshop-api/app/helper/gormhelper"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func (p *PostgreSQL) UpdateOrCreate(data *user_accesses.UserAccess, condition fiber.Map) *gorm.DB {
	var result *gorm.DB
	updateResult := p.DB.Where(condition).Updates(data)
	if gormhelper.HasAffectedRows(updateResult) {
		result = updateResult
	} else {
		result = p.DB.Create(data)
	}
	return result
}
