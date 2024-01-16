package categories

import (
	"time"

	"gopkg.in/guregu/null.v3"
	"gorm.io/gorm"
)

type Category struct {
	ID           int64          `gorm:"column:id;primaryKey;autoIncrement;<-:create"`
	ParentID     null.Int       `gorm:"column:parent_id"`
	Code         string         `gorm:"column:code"`
	Name         string         `gorm:"column:name"`
	IsActive     bool           `gorm:"column:is_active"`
	Icon         null.String    `gorm:"column:icon"`
	DisplayOrder int64          `gorm:"column:display_order"`
	CreatedAt    time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at"`
	Parent       *Category      `gorm:"foreignKey:parent_id;references:id"`
}
