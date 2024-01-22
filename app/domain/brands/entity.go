package brands

import (
	"time"

	"gopkg.in/guregu/null.v3"
	"gorm.io/gorm"
)

type Brand struct {
	ID        int64          `gorm:"column:id;primaryKey;autoIncrement;<-:create"`
	Code      string         `gorm:"column:code;unique"`
	Name      string         `gorm:"column:name"`
	IsActive  bool           `gorm:"column:is_active"`
	Icon      null.String    `gorm:"column:icon"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}
