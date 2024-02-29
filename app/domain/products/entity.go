package products

import (
	"time"

	"gopkg.in/guregu/null.v3"
	"gorm.io/gorm"
)

type Product struct {
	ID                int64          `gorm:"column:id"`
	StoreID           int64          `gorm:"column:store_id"`
	BrandID           int64          `gorm:"column:brand_id"`
	Slug              string         `gorm:"column:slug"`
	Name              string         `gorm:"column:name"`
	IsActive          bool           `gorm:"column:is_active"`
	Price             int64          `gorm:"column:price"`
	Description       null.String    `gorm:"column:description"`
	MinimumPurchase   int64          `gorm:"column:minimum_purchase"`
	VarianCode        string         `gorm:"column:varian_code"`
	Sku               string         `gorm:"column:sku"`
	HasMultipleVarian bool           `gorm:"column:has_multiple_varian"`
	ShortDescription  null.String    `gorm:"column:short_description"`
	Weight            int64          `gorm:"column:weight"`
	Quantity          int64          `gorm:"column:quantity"`
	VirtualQuantity   int64          `gorm:"column:virtual_quantity"`
	CreatedAt         time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt         time.Time      `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	DeletedAt         gorm.DeletedAt `gorm:"column:deleted_at"`
}
