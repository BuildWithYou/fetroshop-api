package users

import "gorm.io/gorm"

type UserRepository interface {
	Create(data *User) *gorm.DB
	Find(destination *User, condition *User) *gorm.DB
}
