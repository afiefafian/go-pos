package model

type CreateSubTotalRequest struct {
	Products []CreateOrderProductRequest `json:"products" validate:"min=1,dive,required"`
}

type CreateOrderProductRequest struct {
	ProductID int64 `json:"productId" validate:"required"`
	Qty       int32 `json:"qty" validate:"required"`
}

type CreateOrderRequest struct {
	PaymentID int64                       `json:"paymentId" validate:"required"`
	TotalPaid int64                       `json:"qty" validate:"required"`
	Products  []CreateOrderProductRequest `json:"products" validate:"required"`
}

type CreateSubTotalResponse struct {
	Subtotal int64                        `json:"subtotal"`
	Products []CreateOrderProductResponse `json:"products"`
}

type CreateOrderProductResponse struct {
	ID               int64                    `json:"productId"`
	Name             string                   `json:"name"`
	Stock            int32                    `json:"stock"`
	Price            int64                    `json:"price"`
	Qty              int32                    `json:"qty"`
	TotalNormalPrice int64                    `json:"totalNormalPrice"`
	TotalFinalPrice  int64                    `json:"totalFinalPrice"`
	Discount         *ProductDiscountResponse `json:"discount"`
}

type CreateOrderResponse struct {
	Order    GetOrderResponse             `json:"order"`
	Products []CreateOrderProductResponse `json:"products"`
}
