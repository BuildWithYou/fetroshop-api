package districts

import (
	"gorm.io/gorm"
)

type DistrictRepo interface {
	Find(destination *District, condition map[string]any) *gorm.DB
	Count(destination *int64, condition map[string]any) *gorm.DB
	List(destination *[]District, condition map[string]any, limit int, offset int, orderBy string) *gorm.DB
}
