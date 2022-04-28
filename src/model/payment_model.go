package model

import "time"

type GetPaymentQuery struct {
	Subtotal   int64
	Pagination PaginationQuery
}

type CreatePaymentRequest struct {
	Name string `json:"name" validate:"required,min=1,max=50"`
	Type string `json:"type" validate:"required,min=1,max=20"`
	Logo string `json:"logo"`
}

type UpdatePaymentRequest struct {
	ID   int64
	Name string `json:"name" validate:"required_without_all=Type Logo,max=50"`
	Type string `json:"type" validate:"required_without_all=Name Logo,max=20"`
	Logo string `json:"logo" validate:"required_without_all=Name Type"`
}

type CreatePaymentResponse struct {
	ID        int64     `json:"paymentId"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	Logo      string    `json:"logo"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type GetPaymentResponse struct {
	ID   int64   `json:"paymentId"`
	Name string  `json:"name"`
	Type string  `json:"type"`
	Logo string  `json:"logo"`
	Card []int64 `json:"card,omitempty"`
}
