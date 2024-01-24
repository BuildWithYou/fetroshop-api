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

type StoreResponse struct {
	Code          string `json:"code"`
	Name          string `json:"name"`
	IsActive      *bool  `json:"isActive"`
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
