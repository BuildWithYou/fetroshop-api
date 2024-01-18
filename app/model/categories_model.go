package model

import (
	"time"

	"gopkg.in/guregu/null.v3"
)

type UpsertCategoryRequest struct {
	ParentCode   string `json:"parentCode" form:"parentCode"`
	Code         string `json:"code" form:"code" validate:"required"`
	Name         string `json:"name" form:"name" validate:"required"`
	IsActive     *bool  `json:"isActive" form:"isActive" validate:"required"`
	Icon         string `json:"icon" form:"icon"`
	DisplayOrder int64  `json:"displayOrder" form:"displayOrder" validate:"required"`
}

type CategoryPathRequest struct {
	Code string `json:"code" form:"code" validate:"required"`
}

type DeleteCategoryRequest struct {
	ForceDelete *bool `json:"forceDelete" form:"forceDelete" validate:"required"`
}

type ListCategoriesRequest struct {
	ParentCode     string `json:"parentCode"`
	Offset         int64  `json:"offset" default:"0"`
	Limit          int64  `json:"limit" default:"10"`
	OrderBy        string `json:"orderBy" validate:"required" enums:"display_order,code,name"`
	OrderDirection string `json:"orderDirection" validate:"required" enums:"ASC,DESC"`
}

type CategoryResponse struct {
	Code         string      ` json:"code"`
	ParentCode   null.String ` json:"parentCode"`
	Name         string      ` json:"name"`
	IsActive     bool        ` json:"isActive"`
	Icon         null.String ` json:"icon"`
	DisplayOrder int64       ` json:"displayOrder"`
	CreatedAt    time.Time   ` json:"createdAt"`
	UpdatedAt    time.Time   ` json:"updatedAt"`
}

type FindCategoryRequest struct {
	Code string `json:"code" validate:"required"`
}
