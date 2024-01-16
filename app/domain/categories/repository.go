package categories

import (
	"gorm.io/gorm"
)

type CategoryRepo interface {
	Create(data *Category) *gorm.DB
	Find(destination *Category, condition map[string]any) *gorm.DB
	List(destination *[]Category, condition map[string]any, limit int, offset int, orderBy string) *gorm.DB
}
