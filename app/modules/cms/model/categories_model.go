package model

type UpsertCategoryRequest struct {
	ParentCode   string `json:"parentCode" validate:"required"`
	Code         string `json:"code" validate:"required"`
	Name         string `json:"name" validate:"required"`
	IsActive     string `json:"isActive"`
	Icon         string `json:"icon"`
	DisplayOrder string `json:"displayOrder" validate:"required"`
}
