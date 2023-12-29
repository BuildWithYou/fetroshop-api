package model

type RegistrationRequest struct {
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	FullName string `json:"fullname"`
}
