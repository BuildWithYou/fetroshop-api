package users

import (
	"gorm.io/gorm"
)

type UserRepo interface {
	Create(data *User) *gorm.DB
	Find(destination *User, condition map[string]any) *gorm.DB
}
