package model

type UpsertStoreRequest struct {
	Code          string `json:"code" validate:"required"` // store code (unique)
	Name          string `json:"name" validate:"required"`
	IsActive      *bool  `json:"isActive" validate:"required"`
	Icon          string `json:"icon"`
	Latitude      string `json:"latitude"`
	Longitude     string `json:"longitude"`
	Address       string `json:"address" validate:"required"`
	ProvinceID    int64  `json:"provinceId" validate:"required"`
	CityID        int64  `json:"cityId" validate:"required"`
	DistrictID    int64  `json:"districtId" validate:"required"`
	SubdistrictID int64  `json:"subdistrictId" validate:"required"`
	PostalCode    string `json:"postalCode" validate:"required"`
}

type StoreListRequest struct {
	Search         string `json:"search"` // Store Code or Store Name
	Offset         int64  `json:"offset" default:"0"`
	Limit          int64  `json:"limit" default:"10"`
	OrderBy        string `json:"orderBy" validate:"required" enums:"code,name,updated_at,created_at"`
	OrderDirection string `json:"orderDirection" validate:"required" enums:"ASC,DESC"`
}

// #marked: for swagger generation purposes only
//
//lint:ignore U1000 Ignore unused code
type storeDetailResponse struct {
	Code    int            `json:"code"`    // http status code
	Status  string         `json:"status"`  // http status message
	Message string         `json:"message"` // message from system
	Data    StoreDetail    `json:"data"`    // main data
	Meta    any            `json:"meta"`    // support data
	Errors  map[string]any `json:"errors"`  // error data
}

// #marked: for swagger generation purposes only
//
//lint:ignore U1000 Ignore unused code
type storesListResponse struct {
	Code    int             `json:"code"`    // http status code
	Status  string          `json:"status"`  // http status message
	Message string          `json:"message"` // message from system
	Data    []StoreListData `json:"data"`    // main data
	Meta    any             `json:"meta"`    // support data
	Errors  map[string]any  `json:"errors"`  // error data
}

type StoreListData struct {
	Code          string  `json:"code"`
	Name          string  `json:"name"`
	IsActive      bool    `json:"isActive"`
	Icon          *string `json:"icon"`
	Latitude      *string `json:"latitude"`
	Longitude     *string `json:"longitude"`
	Address       string  `json:"address"`
	ProvinceID    int64   `json:"provinceId"`
	CityID        int64   `json:"cityId"`
	DistrictID    int64   `json:"districtId"`
	SubdistrictID int64   `json:"subdistrictId"`
	PostalCode    string  `json:"postalCode"`
}

type StoreDetail struct {
	Code        string   `json:"code"`
	Name        string   `json:"name"`
	IsActive    bool     `json:"isActive"`
	Icon        *string  `json:"icon"`
	Latitude    *string  `json:"latitude"`
	Longitude   *string  `json:"longitude"`
	Address     string   `json:"address"`
	Province    Location `json:"province"`
	City        Location `json:"city"`
	District    Location `json:"district"`
	Subdistrict Location `json:"subdistrict"`
	PostalCode  string   `json:"postalCode"`
}
