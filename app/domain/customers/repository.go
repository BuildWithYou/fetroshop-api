package customers

import "gorm.io/gorm"

type CustomerRepo interface {
	Create(data *Customer) *gorm.DB
	Find(destination *Customer, condition map[string]any) *gorm.DB
}
