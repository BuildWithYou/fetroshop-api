package brands

import (
	"gorm.io/gorm"
)

type BrandRepo interface {
	Create(data *Brand) *gorm.DB
	Find(destination *Brand, condition map[string]any) *gorm.DB
	List(destination *[]Brand, condition map[string]any, limit int, offset int, orderBy string) *gorm.DB
	Update(data *Brand, condition map[string]any) *gorm.DB
	Delete(condition map[string]any) *gorm.DB
}
