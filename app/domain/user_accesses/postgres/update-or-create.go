package postgres

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/user_accesses"
	"github.com/BuildWithYou/fetroshop-api/app/helper/gormhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
	"gorm.io/gorm"
)

func (p *PostgreSQL) UpdateOrCreate(data *user_accesses.UserAccess, condition *user_accesses.UserAccess) *gorm.DB {
	var result *gorm.DB
	existingData := new(user_accesses.UserAccess)
	searchResult := p.DB.Where(condition).First(existingData, condition)
	if validatorhelper.IsNotNil(searchResult.Error) && !gormhelper.IsErrRecordNotFound(searchResult.Error) {
		return searchResult
	}
	if gormhelper.IsErrRecordNotFound(searchResult.Error) {
		result = p.DB.Create(data)
	} else {
		result = p.DB.Where(existingData).Updates(data)
	}
	return result

	/* var result *gorm.DB
	updateResult := p.DB.Where(condition).Updates(data)
	if gormhelper.HasAffectedRows(updateResult) {
		result = updateResult
	} else {
		result = p.DB.Create(data)
	}
	return result */
}
