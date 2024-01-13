package customer_accesses

import (
	"gorm.io/gorm"
)

type CustomerAccessRepo interface {
	Create(data *CustomerAccess) *gorm.DB
	UpdateOrCreate(data *CustomerAccess, condition map[string]any) *gorm.DB
	Update(data *CustomerAccess, condition map[string]any) *gorm.DB
	Find(destination *CustomerAccess, condition map[string]any) *gorm.DB
	Delete(condition *CustomerAccess) *gorm.DB
}
