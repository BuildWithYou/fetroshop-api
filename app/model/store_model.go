package model

type UpsertStoreRequest struct {
	Code          string `json:"code" validate:"required"`
	Name          string `json:"name" validate:"required"`
	IsActive      *bool  `json:"isActive" validate:"required"`
	Icon          string `json:"icon"`
	Latitude      string `json:"latitude"`
	Longitude     string `json:"longitude"`
	Address       string `json:"address"`
	ProvinceID    string `json:"provinceId"`
	CityID        string `json:"cityId"`
	DistrictID    string `json:"districtId"`
	SubdistrictID string `json:"subdistrictId"`
	PostalCode    string `json:"postalCode"`
}

type ListStoresRequest struct {
	Offset         int64  `json:"offset" default:"0"`
	Limit          int64  `json:"limit" default:"10"`
	OrderBy        string `json:"orderBy" validate:"required" enums:"code,name,updated_at,created_at"`
	OrderDirection string `json:"orderDirection" validate:"required" enums:"ASC,DESC"`
}

// #marked: for swagger generation purposes only
type StoreDetailResponse struct {
	Code    int            `json:"code"`    // http status code
	Status  string         `json:"status"`  // http status message
	Message string         `json:"message"` // message from system
	Data    StoreDetail    `json:"data"`    // main data
	Meta    any            `json:"meta"`    // support data
	Errors  map[string]any `json:"errors"`  // error data
}

// #marked: for swagger generation purposes only
type StoresListResponse struct {
	Code    int            `json:"code"`    // http status code
	Status  string         `json:"status"`  // http status message
	Message string         `json:"message"` // message from system
	Data    []StoreDetail  `json:"data"`    // main data
	Meta    any            `json:"meta"`    // support data
	Errors  map[string]any `json:"errors"`  // error data
}

type StoreDetail struct {
	Code        string   `json:"code"`
	Name        string   `json:"name"`
	IsActive    *bool    `json:"isActive"`
	Icon        string   `json:"icon"`
	Latitude    string   `json:"latitude"`
	Longitude   string   `json:"longitude"`
	Address     string   `json:"address"`
	Province    Location `json:"province"`
	City        Location `json:"city"`
	District    Location `json:"district"`
	Subdistrict Location `json:"subdistrict"`
	PostalCode  string   `json:"postalCode"`
}

type Location struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
