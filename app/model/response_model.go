package model

type Response struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}
