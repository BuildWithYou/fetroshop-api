package customers

import "gorm.io/gorm"

type CustomerRepository interface {
	Create(cst *Customer) *gorm.DB
	Find(destination *Customer, condition *Customer) *gorm.DB
}
