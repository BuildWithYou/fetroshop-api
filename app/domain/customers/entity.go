package customers

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	ID        int64          `gorm:"column:id;primaryKey;autoIncrement;<-:create" json:"id"`
	Username  string         `gorm:"column:username" json:"username"`
	Phone     string         `gorm:"column:phone" json:"phone"`
	Email     string         `gorm:"column:email" json:"email"`
	FullName  string         `gorm:"column:full_name" json:"fullName"`
	Password  string         `gorm:"column:password" json:"password"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoCreateTime;autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deletedAt"`
}
