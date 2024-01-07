package user_accesses

import (
	"gorm.io/gorm"
)

type UserAccessRepo interface {
	Create(data *UserAccess) *gorm.DB
	Find(destination *UserAccess, condition *UserAccess) *gorm.DB
}
