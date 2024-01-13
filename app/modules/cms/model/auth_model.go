package model

type RegistrationRequest struct {
	Username string `json:"username" form:"username" validate:"required"`
	Phone    string `json:"phone" form:"phone" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	FullName string `json:"fullName" form:"fullName" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

type LoginRequest struct {
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}
