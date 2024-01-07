package customer_accesses

import (
	"gorm.io/gorm"
)

type CustomerAccessRepo interface {
	Create(data *CustomerAccess) *gorm.DB
	UpdateOrCreate(data *CustomerAccess, condition *CustomerAccess) *gorm.DB
	Find(destination *CustomerAccess, condition *CustomerAccess) *gorm.DB
}
