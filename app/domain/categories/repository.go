package categories

import "gorm.io/gorm"

type CategoryRepo interface {
	Create(data *Category) *gorm.DB
	Find(destination *Category, condition *Category) *gorm.DB
}
