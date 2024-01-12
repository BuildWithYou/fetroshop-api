package model

type ListCategoriesRequest struct {
	Offset         int64  `json:"offset" validate:"required" default:"0"`
	Limit          int64  `json:"limit" validate:"required" default:"10"`
	OrderBy        string `json:"orderBy" validate:"required" enums:"displayOrder,code,name"`
	OrderDirection string `json:"orderDirection" validate:"required" enums:"ASC,DESC"`
}

type FindCategoryRequest struct {
	Code string `json:"code" validate:"required"`
}
