package postgres

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/customer_accesses"
	"github.com/BuildWithYou/fetroshop-api/app/helper/gormhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
	"gorm.io/gorm"
)

func (p *PostgreSQL) UpdateOrCreate(data *customer_accesses.CustomerAccess, condition *customer_accesses.CustomerAccess) *gorm.DB {
	var result *gorm.DB
	existingData := new(customer_accesses.CustomerAccess)
	searchResult := p.DB.Where(condition).First(existingData, condition)
	if validatorhelper.IsNotNil(searchResult.Error) && !gormhelper.IsRecordNotFound(searchResult.Error) {
		return searchResult
	}
	if !gormhelper.IsRecordNotFound(searchResult.Error) {
		// result = p.DB.Where(existingData).Updates(data)
		result = p.DB.Where(data).Updates(existingData)
	} else {
		result = p.DB.Create(data)
	}
	return result
}
