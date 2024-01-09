package postgres

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/customer_accesses"
	"github.com/BuildWithYou/fetroshop-api/app/helper/gormhelper"
	"gorm.io/gorm"
)

/*
#marked: buggy
It failed update because it's id included even it doesn't exist is condition arg
*/
func (p *PostgreSQL) UpdateOrCreate(data *customer_accesses.CustomerAccess, condition *customer_accesses.CustomerAccess) *gorm.DB {
	var result *gorm.DB
	updateResult := p.DB.Where(condition).Updates(data)
	if gormhelper.HasAffectedRows(updateResult) {
		result = updateResult
	} else {
		result = p.DB.Create(data)
	}
	return result
}
