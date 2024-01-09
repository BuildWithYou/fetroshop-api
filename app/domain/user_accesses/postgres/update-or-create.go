package postgres

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/user_accesses"
	"github.com/BuildWithYou/fetroshop-api/app/helper/gormhelper"
	"gorm.io/gorm"
)

func (p *PostgreSQL) UpdateOrCreate(data *user_accesses.UserAccess, condition *user_accesses.UserAccess) *gorm.DB {
	/* var result *gorm.DB
	existingData := new(user_accesses.UserAccess)
	searchResult := p.DB.Where(condition).First(existingData, condition)
	if validatorhelper.IsNotNil(searchResult.Error) && !gormhelper.IsRecordNotFound(searchResult.Error) {
		return searchResult
	}
	if !gormhelper.IsRecordNotFound(searchResult.Error) {
		result = p.DB.Where(existingData).Updates(data)
	} else {
		result = p.DB.Create(data)
	}
	return result */

	var result *gorm.DB
	updateResult := p.DB.Where(condition).Updates(data)
	if gormhelper.HasAffectedRows(updateResult) {
		result = updateResult
	} else {
		result = p.DB.Create(data)
	}
	return result
}
