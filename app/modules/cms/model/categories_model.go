package model

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
