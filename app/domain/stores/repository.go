package stores

import (
	"gorm.io/gorm"
)

type StoreRepo interface {
	Create(data *Store) *gorm.DB
	Find(destination *Store, condition map[string]any) *gorm.DB
	List(destination *[]Store, condition map[string]any, limit int, offset int, orderBy string) *gorm.DB
	Update(data *Store, condition map[string]any) *gorm.DB
	Delete(condition map[string]any) *gorm.DB
}
