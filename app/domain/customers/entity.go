package customers

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	ID        int64          `gorm:"column:id;primaryKey;autoIncrement;<-:create"`
	Username  int64          `gorm:"column:username"`
	Phone     string         `gorm:"column:phone"`
	Email     string         `gorm:"column:email"`
	FullName  string         `gorm:"column:full_name"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}
