package model

import "time"

type ProvinceListRequest struct {
	Name           string `json:"name"`
	Offset         int64  `json:"offset" default:"0"`
	Limit          int64  `json:"limit" default:"10"`
	OrderBy        string `json:"orderBy" validate:"required" enums:"id,name,updated_at,created_at"`
	OrderDirection string `json:"orderDirection" validate:"required" enums:"ASC,DESC"`
}

// #marked: for swagger generation purposes only
type ProvinceListResponse struct {
	Code    int              `json:"code"`    // http status code
	Status  string           `json:"status"`  // http status message
	Message string           `json:"message"` // message from system
	Data    []ProvinceDetail `json:"data"`    // main data
	Meta    listMeta         `json:"meta"`    // support data
	Errors  map[string]any   `json:"errors"`  // error data
}

type ProvinceDetail struct {
	ID        int64     ` json:"id"`
	Name      string    ` json:"name"`
	CreatedAt time.Time ` json:"createdAt"`
	UpdatedAt time.Time ` json:"updatedAt"`
}
