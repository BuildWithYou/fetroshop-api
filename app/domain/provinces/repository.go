package provinces

import (
	"gorm.io/gorm"
)

type ProvinceRepo interface {
	Find(destination *Province, condition map[string]any) *gorm.DB
	Count(destination *int64, condition map[string]any) *gorm.DB
	List(destination *[]Province, condition map[string]any, limit int, offset int, orderBy string) *gorm.DB
}
