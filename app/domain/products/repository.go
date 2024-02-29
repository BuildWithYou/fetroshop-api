package products

import (
	"gorm.io/gorm"
)

type ProductRepo interface {
	Create(data *Product) *gorm.DB
	Find(destination *Product, condition map[string]any) *gorm.DB
	List(destination *[]Product, keyword string, limit int, offset int, orderBy string) *gorm.DB
	Update(data *Product, condition map[string]any) *gorm.DB
	Delete(condition map[string]any) *gorm.DB
	Count(destination *int64, keyword string) *gorm.DB
}
