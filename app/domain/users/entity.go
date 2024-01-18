package users

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int64          `gorm:"column:id;primaryKey;autoIncrement;<-:create" json:"id"`
	Username  string         `gorm:"column:username;unique" json:"username"`
	Phone     string         `gorm:"column:phone;unique" json:"phone"`
	Email     string         `gorm:"column:email;unique" json:"email"`
	FullName  string         `gorm:"column:full_name" json:"fullName"`
	Password  string         `gorm:"column:password"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoCreateTime;autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deletedAt"`
}
