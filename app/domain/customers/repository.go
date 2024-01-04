package customers

import "gorm.io/gorm"

type CustomerRepository interface {
	Create(cst *Customer) *gorm.DB
	Find(cst *Customer) *Customer
}
