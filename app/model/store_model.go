package model

import (
	"time"

	"gopkg.in/guregu/null.v3"
)

type UpsertStoreRequest struct {
	ParentCode   string `json:"parentCode"`
	Code         string `json:"code" validate:"required"`
	Name         string `json:"name" validate:"required"`
	IsActive     *bool  `json:"isActive" validate:"required"`
	Icon         string `json:"icon"`
	DisplayOrder int64  `json:"displayOrder" validate:"required"`
}

type ListStoresRequest struct {
	ParentCode     string `json:"parentCode"`
	Offset         int64  `json:"offset" default:"0"`
	Limit          int64  `json:"limit" default:"10"`
	OrderBy        string `json:"orderBy" validate:"required" enums:"display_order,code,name,updated_at,created_at"`
	OrderDirection string `json:"orderDirection" validate:"required" enums:"ASC,DESC"`
}

type StoreResponse struct {
	Code         string      ` json:"code"`
	ParentCode   null.String ` json:"parentCode"`
	Name         string      ` json:"name"`
	IsActive     bool        ` json:"isActive"`
	Icon         null.String ` json:"icon"`
	DisplayOrder int64       ` json:"displayOrder"`
	CreatedAt    time.Time   ` json:"createdAt"`
	UpdatedAt    time.Time   ` json:"updatedAt"`
}
