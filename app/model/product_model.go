package model

type UpsertProductRequest struct {
	StoreCode        string   `form:"storeCode" validate:"required"`
	BrandCode        string   `form:"brandCode" validate:"required"`
	Name             string   `form:"name" validate:"required"`
	IsActive         bool     `form:"isActive" validate:"required,boolean"`
	Price            int64    `form:"price" validate:"required,number"`
	Description      string   `form:"description" validate:"required"`
	MinimumPurchase  int64    `form:"minimumPurchase" validate:"required,number"`
	VarianCode       string   `form:"varianCode" validate:"required"`
	Sku              string   `form:"sku" validate:"required"`
	ShortDescription string   `form:"shortDescription" validate:"required"`
	Weight           int64    `form:"weight" validate:"required,number"` // gram
	Quantity         int64    `form:"quantity" validate:"required,number"`
	ImageUrl         []string `form:"imageUrl"`
	VideoUrl         []string `form:"videoUrl"`
}

type ProductListRequest struct {
	Search         string `json:"search"` // Product Code or Product Name
	Offset         int64  `json:"offset" default:"0"`
	Limit          int64  `json:"limit" default:"10"`
	OrderBy        string `json:"orderBy" validate:"required" enums:"code,name,updated_at,created_at"`
	OrderDirection string `json:"orderDirection" validate:"required" enums:"ASC,DESC"`
}

// #marked: for swagger generation purposes only
//
//lint:ignore U1000 Ignore unused code
type productDetailResponse struct {
	Code    int            `json:"code"`    // http status code
	Status  string         `json:"status"`  // http status message
	Message string         `json:"message"` // message from system
	Data    ProductDetail  `json:"data"`    // main data
	Meta    any            `json:"meta"`    // support data
	Errors  map[string]any `json:"errors"`  // error data
}

// #marked: for swagger generation purposes only
//
//lint:ignore U1000 Ignore unused code
type productsListResponse struct {
	Code    int               `json:"code"`    // http status code
	Status  string            `json:"status"`  // http status message
	Message string            `json:"message"` // message from system
	Data    []ProductListData `json:"data"`    // main data
	Meta    any               `json:"meta"`    // support data
	Errors  map[string]any    `json:"errors"`  // error data
}

type ProductListData struct {
	StoreID           int64   `json:"storeId"`
	BrandID           int64   `json:"brandId"`
	Slug              string  `json:"slug"`
	Name              string  `json:"name"`
	IsActive          bool    `json:"isActive"`
	Price             int64   `json:"price"`
	Description       *string `json:"description"`
	MinimumPurchase   int64   `json:"minimumPurchase"`
	VarianCode        string  `json:"varianCode"`
	Sku               string  `json:"sku"`
	HasMultipleVarian bool    `json:"hasMultipleVarian"`
	ShortDescription  *string `json:"shortDescription"`
	Weight            int64   `json:"weight"`
	Quantity          int64   `json:"quantity"`
	VirtualQuantity   int64   `json:"virtualQuantity"`
}

type ProductDetail struct {
	StoreID           int64   `json:"storeId"`
	BrandID           int64   `json:"brandId"`
	Slug              string  `json:"slug"`
	Name              string  `json:"name"`
	IsActive          bool    `json:"isActive"`
	Price             int64   `json:"price"`
	Description       *string `json:"description"`
	MinimumPurchase   int64   `json:"minimumPurchase"`
	VarianCode        string  `json:"varianCode"`
	Sku               string  `json:"sku"`
	HasMultipleVarian bool    `json:"hasMultipleVarian"`
	ShortDescription  *string `json:"shortDescription"`
	Weight            int64   `json:"weight"`
	Quantity          int64   `json:"quantity"`
	VirtualQuantity   int64   `json:"virtualQuantity"`
	Store             IDName  `json:"store"`
	Brand             IDName  `json:"brand"`
}
