package model

type CmsRegistrationRequest struct {
	Username string `json:"username" form:"username" validate:"required"`
	Phone    string `json:"phone" form:"phone" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	FullName string `json:"fullName" form:"fullName" validate:"required"`
	Password string `json:"password" form:"password" validate:"required,min=8"`
}

type CmsLoginRequest struct {
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

type WebRegistrationRequest struct {
	Username string `json:"username" form:"username" validate:"required"`
	Phone    string `json:"phone" form:"phone" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	FullName string `json:"fullName" form:"fullName" validate:"required"`
	Password string `json:"password" form:"password" validate:"required,min=8"`
}

type WebLoginRequest struct {
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}
