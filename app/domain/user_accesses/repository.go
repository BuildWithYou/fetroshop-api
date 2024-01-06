package user_accesses

import (
	"gorm.io/gorm"
)

type UserAccessRepository interface {
	Create(data *UserAccess) *gorm.DB
	Find(destination *UserAccess, condition *UserAccess) *gorm.DB
}
