package categories

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID           int64          `gorm:"column:id;primaryKey;autoIncrement;<-:create" json:"id"`
	ParentID     int64          `gorm:"column:parent_id" json:"parentID"`
	Code         string         `gorm:"column:code" json:"code"`
	Name         string         `gorm:"column:name" json:"name"`
	IsActive     bool           `gorm:"column:is_active" json:"isActive"`
	Icon         string         `gorm:"column:icon" json:"icon"`
	DisplayOrder int64          `gorm:"column:display_order" json:"displayOrder"`
	CreatedAt    time.Time      `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;autoCreateTime;autoUpdateTime" json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at" json:"deletedAt"`
}
