package provinces

import (
	"time"

	"gorm.io/gorm"
)

type Province struct {
	ID        int64          `gorm:"column:id;primaryKey;autoIncrement;<-:create"`
	Name      string         `gorm:"column:name"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}
