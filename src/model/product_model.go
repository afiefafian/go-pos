package model

import "time"

type GetProductQuery struct {
	CategoryId *int64
	Q          *string
	Pagination PaginationQuery
}

type CreateProductRequest struct {
	CategoryID int64                         `json:"categoryId" validate:"required"`
	Name       string                        `json:"name" validate:"required,min=1"`
	Image      string                        `json:"image" validate:"required,min=1"`
	Stock      int32                         `json:"stock" validate:"required"`
	Price      int64                         `json:"price" validate:"required"`
	Discount   *CreateProductDiscountRequest `json:"discount"`
}

type CreateProductDiscountRequest struct {
	Type      string `json:"type" validate:"required,min=1"`
	Qty       int32  `json:"qty" validate:"required"`
	Result    int64  `json:"result" validate:"required"`
	ExpiredAt int64  `json:"expiredAt" validate:"required"`
}

type UpdateProductRequest struct {
	ID         int64
	CategoryID int64  `json:"categoryId" validate:"required"`
	Name       string `json:"name" validate:"required,min=1"`
	Image      string `json:"image" validate:"required,min=1"`
	Stock      int32  `json:"stock" validate:"required"`
	Price      int64  `json:"price" validate:"required"`
}

type GetProductResponse struct {
	ID       int64                    `json:"productId"`
	SKU      string                   `json:"sku"`
	Name     string                   `json:"name"`
	Stock    int32                    `json:"stock"`
	Price    int64                    `json:"price"`
	Image    string                   `json:"image"`
	Category *GetCategoryResponse     `json:"category"`
	Discount *ProductDiscountResponse `json:"discount"`
}

type ProductDiscountResponse struct {
	ID              int64      `json:"discountId"`
	Type            string     `json:"type"`
	Qty             int32      `json:"qty"`
	Result          int64      `json:"result"`
	ExpiredAt       *time.Time `json:"expiredAt"`
	ExpiredAtFormat string     `json:"expiredAtFormat"`
	StringFormat    string     `json:"stringFormat"`
}

type CreateProductResponse struct {
	ID         int64     `json:"productId"`
	CategoryID int64     `json:"categoryId"`
	SKU        string    `json:"sku"`
	Name       string    `json:"name"`
	Stock      int32     `json:"stock"`
	Price      int64     `json:"price"`
	Image      string    `json:"image"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
