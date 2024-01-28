package subdistricts

import (
	"gorm.io/gorm"
)

type SubdistrictRepo interface {
	Find(destination *Subdistrict, condition map[string]any) *gorm.DB
	Count(destination *int64, condition map[string]any) *gorm.DB
	List(destination *[]Subdistrict, condition map[string]any, limit int, offset int, orderBy string) *gorm.DB
}
