package model

type UpsertCategoryRequest struct {
	ParentCode   string `json:"parentCode"`
	Code         string `json:"code" validate:"required"`
	Name         string `json:"name" validate:"required"`
	IsActive     bool   `json:"isActive"`
	Icon         string `json:"icon"`
	DisplayOrder int64  `json:"displayOrder" validate:"required"`
}
