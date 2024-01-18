package model

type RegistrationRequest struct {
	Username string `json:"username" form:"username" validate:"required"`
	Phone    string `json:"phone" form:"phone" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	FullName string `json:"fullName" form:"fullName" validate:"required"`
	Password string `json:"password" form:"password" validate:"required,min=8"`
}

type LoginRequest struct {
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

type RegistrationResponseData struct {
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	FullName string `json:"fullName"`
}
