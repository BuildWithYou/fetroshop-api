package districts

import (
	"time"

	"gorm.io/gorm"
)

type District struct {
	ID        int64          `gorm:"column:id;primaryKey;autoIncrement;<-:create"`
	CityID    int64          `gorm:"column:city_id"`
	Name      string         `gorm:"column:name"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}
