package model

type PaginationQuery struct {
	Limit int8
	Skip  int8
}

type BaseResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

type PaginationResponse struct {
	Total int8 `json:"total"`
	Limit int8 `json:"limit"`
	Skip  int8 `json:"skip"`
}
