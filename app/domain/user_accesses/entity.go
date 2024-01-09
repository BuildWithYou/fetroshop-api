package user_accesses

import (
	"time"

	"gorm.io/gorm"
)

type UserAccess struct {
	ID        int64          `gorm:"column:id;primaryKey;autoIncrement;" json:"id"`
	UserID    int64          `gorm:"column:user_id" json:"userId"`
	Key       string         `gorm:"column:key" json:"key"`
	Platform  string         `gorm:"column:platform" json:"platform"`
	UserAgent string         `gorm:"column:user_agent" json:"userAgent"`
	ExpiredAt time.Time      `gorm:"column:expired_at;" json:"expiredAt"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoCreateTime;autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deletedAt"`
}
