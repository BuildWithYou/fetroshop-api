package user_accesses

import (
	"gorm.io/gorm"
)

type UserAccessRepo interface {
	Create(data *UserAccess) *gorm.DB
	Find(destination *UserAccess, condition *UserAccess) *gorm.DB
	UpdateOrCreate(data *UserAccess, condition *UserAccess) *gorm.DB
	Delete(condition *UserAccess) *gorm.DB
	Update(data *UserAccess, condition *UserAccess) *gorm.DB
}
