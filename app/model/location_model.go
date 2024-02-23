package model

// Province
type ProvinceListRequest struct {
	Name           string `json:"name"`
	Offset         int64  `json:"offset" default:"0"`
	Limit          int64  `json:"limit" default:"10"`
	OrderBy        string `json:"orderBy" validate:"required" enums:"id,name,updated_at,created_at"`
	OrderDirection string `json:"orderDirection" validate:"required" enums:"ASC,DESC"`
}

// #marked: for swagger generation purposes only
//
//lint:ignore U1000 Ignore unused code
type locationListResponse struct {
	Code    int            `json:"code"`    // http status code
	Status  string         `json:"status"`  // http status message
	Message string         `json:"message"` // message from system
	Data    []IDName       `json:"data"`    // main data
	Meta    listMeta       `json:"meta"`    // support data
	Errors  map[string]any `json:"errors"`  // error data
}

// City
type CityListRequest struct {
	ProvinceID     string `json:"provinceId" validate:"required"`
	Name           string `json:"name"`
	Offset         int64  `json:"offset" default:"0"`
	Limit          int64  `json:"limit" default:"10"`
	OrderBy        string `json:"orderBy" validate:"required" enums:"id,name,updated_at,created_at"`
	OrderDirection string `json:"orderDirection" validate:"required" enums:"ASC,DESC"`
}

// District
type DistrictListRequest struct {
	CityID         string `json:"cityId" validate:"required"`
	Name           string `json:"name"`
	Offset         int64  `json:"offset" default:"0"`
	Limit          int64  `json:"limit" default:"10"`
	OrderBy        string `json:"orderBy" validate:"required" enums:"id,name,updated_at,created_at"`
	OrderDirection string `json:"orderDirection" validate:"required" enums:"ASC,DESC"`
}

// Subdistrict
type SubdistrictListRequest struct {
	DistrictID     string `json:"districtId" validate:"required"`
	Name           string `json:"name"`
	Offset         int64  `json:"offset" default:"0"`
	Limit          int64  `json:"limit" default:"10"`
	OrderBy        string `json:"orderBy" validate:"required" enums:"id,name,updated_at,created_at"`
	OrderDirection string `json:"orderDirection" validate:"required" enums:"ASC,DESC"`
}
