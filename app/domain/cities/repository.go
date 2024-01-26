package cities

import (
	"gorm.io/gorm"
)

type CityRepo interface {
	Find(destination *City, condition map[string]any) *gorm.DB
	List(destination *[]City, condition map[string]any, limit int, offset int, orderBy string) *gorm.DB
}
