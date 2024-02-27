package product_medias

import (
	"time"

	"gorm.io/gorm"
)

type ProductMedia struct {
	ID        int64          `gorm:"column:id"`
	ProductID int64          `gorm:"column:product_id"`
	File      string         `gorm:"column:file"`
	MediaType string         `gorm:"column:media_type"` // Available options: image_file, image_url, video_file, video_url
	Size      string         `gorm:"column:is_active"`  // Available options: S, M, L, XL
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}
