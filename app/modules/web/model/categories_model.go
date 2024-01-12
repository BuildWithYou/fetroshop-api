package model

type ListCategoriesRequest struct {
	Username string `json:"username" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
	Email    string `json:"email" validate:"required"`
	FullName string `json:"fullName" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type FindCategoryRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
