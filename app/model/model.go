package model

type Response struct {
	Code    int            `json:"code"`    // http status code
	Status  string         `json:"status"`  // http status message
	Message string         `json:"message"` // message from system
	Data    any            `json:"data"`    // main data
	Meta    any            `json:"meta"`    // support data
	Errors  map[string]any `json:"errors"`  // error data
}

type FindByCodeRequest struct {
	Code string `json:"code" validate:"required"`
}

type DeleteRequest struct {
	ForceDelete *bool `json:"forceDelete" validate:"required"`
}
