package model

import "time"

type GetPaymentQuery struct {
	subtotal   int64
	pagination PaginationQuery
}

type CreatePaymentRequest struct {
	Name string `json:"name" validate:"required,min=1,max=50"`
	Type string `json:"type" validate:"required,min=1,max=20"`
	Logo string `json:"logo"`
}

type UpdatePaymentRequest struct {
	ID   int64
	Name string `json:"name" validate:"required,min=1,max=50"`
	Type string `json:"type" validate:"required,min=1,max=20"`
	Logo string `json:"logo"`
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
	ID   int64  `json:"paymentId"`
	Name string `json:"name"`
	Type string `json:"type"`
	Logo string `json:"logo"`
}
