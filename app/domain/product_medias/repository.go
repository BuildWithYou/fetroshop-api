package product_medias

import (
	"gorm.io/gorm"
)

type ProductMediaRepo interface {
	Create(data *ProductMedia) *gorm.DB
	Find(destination *ProductMedia, condition map[string]any) *gorm.DB
	Update(data *ProductMedia, condition map[string]any) *gorm.DB
	Delete(condition map[string]any) *gorm.DB
}
