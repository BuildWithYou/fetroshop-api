package model

type ListCategoriesRequest struct {
	Offset         int64 `json:"offset" validate:"required"`
	Limit          int64 `json:"limit" validate:"required"`
	OrderBy        int64 `json:"orderBy" validate:"required"`
	OrderDirection int64 `json:"orderDirection" validate:"required"`
}

type FindCategoryRequest struct {
	Code string `json:"code" validate:"required"`
}
