package customer_accesses

import (
	"gorm.io/gorm"
)

type CustomerAccessRepository interface {
	Create(data *CustomerAccess) *gorm.DB
	Find(destination *CustomerAccess, condition *CustomerAccess) *gorm.DB
}
