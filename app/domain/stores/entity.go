package stores

import (
	"time"

	"github.com/BuildWithYou/fetroshop-api/app/domain/cities"
	"github.com/BuildWithYou/fetroshop-api/app/domain/districts"
	"github.com/BuildWithYou/fetroshop-api/app/domain/provinces"
	"github.com/BuildWithYou/fetroshop-api/app/domain/subdistricts"
	"gopkg.in/guregu/null.v3"
	"gorm.io/gorm"
)

type Store struct {
	ID            int64                     `gorm:"column:id;primaryKey;autoIncrement;<-:create"`
	UserID        int64                     `gorm:"column:user_id"`
	Code          string                    `gorm:"column:code;unique"`
	Name          string                    `gorm:"column:name"`
	IsActive      bool                      `gorm:"column:is_active"`
	Icon          null.String               `gorm:"column:icon"`
	Latitude      null.String               `gorm:"column:latitude"`
	Longitude     null.String               `gorm:"column:longitude"`
	Address       string                    `gorm:"column:address"`
	ProvinceID    int64                     `gorm:"column:province_id"`
	CityID        int64                     `gorm:"column:city_id"`
	DistrictID    int64                     `gorm:"column:district_id"`
	SubdistrictID int64                     `gorm:"column:subdistrict_id"`
	PostalCode    string                    `gorm:"column:postal_code"`
	CreatedAt     time.Time                 `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt     time.Time                 `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	DeletedAt     gorm.DeletedAt            `gorm:"column:deleted_at"`
	Province      *provinces.Province       `gorm:"foreignKey:province_id;references:id"`
	City          *cities.City              `gorm:"foreignKey:city_id;references:id"`
	District      *districts.District       `gorm:"foreignKey:district_id;references:id"`
	Subdistrict   *subdistricts.Subdistrict `gorm:"foreignKey:subdistrict_id;references:id"`
}
