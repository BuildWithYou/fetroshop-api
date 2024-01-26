package stores

import (
	"time"

	"gopkg.in/guregu/null.v3"
	"gorm.io/gorm"
)

type Store struct {
	ID            int64          `gorm:"column:id;primaryKey;autoIncrement;<-:create"`
	UserID        int64          `gorm:"column:parent_id"`
	Code          string         `gorm:"column:code;unique"`
	Name          string         `gorm:"column:name"`
	IsActive      bool           `gorm:"column:is_active"`
	Icon          null.String    `gorm:"column:icon"`
	Latitude      string         `gorm:"column:latitude"`
	Longitude     string         `gorm:"column:longitude"`
	Address       string         `gorm:"column:address"`
	ProvinceID    int64          `gorm:"column:province_id"`
	CityID        int64          `gorm:"column:city_id"`
	DistrictID    int64          `gorm:"column:district_id"`
	SubdistrictID int64          `gorm:"column:subdistrict_id"`
	PostalCode    string         `gorm:"column:postal_code"`
	CreatedAt     time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt     time.Time      `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at"`
}
