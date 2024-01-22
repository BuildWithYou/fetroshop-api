package model

import (
	"time"

	"gopkg.in/guregu/null.v3"
)

type UpsertBrandRequest struct {
	Code     string `json:"code" validate:"required"`
	Name     string `json:"name" validate:"required"`
	IsActive *bool  `json:"isActive" validate:"required"`
	Icon     string `json:"icon"`
}

type ListBrandsRequest struct {
	Offset         int64  `json:"offset" default:"0"`
	Limit          int64  `json:"limit" default:"10"`
	OrderBy        string `json:"orderBy" validate:"required" enums:"code,name,updated_at,created_at"`
	OrderDirection string `json:"orderDirection" validate:"required" enums:"ASC,DESC"`
}

type ListBrandsByPrefixRequest struct {
	Prefix         string `json:"prefix" validate:"required"`
	Offset         int64  `json:"offset" default:"0"`
	Limit          int64  `json:"limit" default:"10"`
	OrderBy        string `json:"orderBy" validate:"required" enums:"code,name,updated_at,created_at"`
	OrderDirection string `json:"orderDirection" validate:"required" enums:"ASC,DESC"`
}

type BrandResponse struct {
	Code      string      ` json:"code"`
	Name      string      ` json:"name"`
	IsActive  bool        ` json:"isActive"`
	Icon      null.String ` json:"icon"`
	CreatedAt time.Time   ` json:"createdAt"`
	UpdatedAt time.Time   ` json:"updatedAt"`
}
