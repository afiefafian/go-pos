package model

import "time"

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
	Order    GetOrderResponse            `json:"order"`
	Products []CreateOrderProductRequest `json:"products" validate:"required"`
}

type GetOrderResponse struct {
	ID            int64                    `json:"orderId"`
	CashierID     int64                    `json:"cashiersId"`
	PaymentTypeID int64                    `json:"paymentTypesId"`
	TotalPrice    int64                    `json:"totalPrice"`
	TotalPaid     int64                    `json:"totalPaid"`
	TotalReturn   int64                    `json:"totalReturn"`
	ReceiptID     string                   `json:"receiptId"`
	CreatedAt     time.Time                `json:"createdAt"`
	Cashier       *GetCashierResponse      `json:"cashier"`
	PaymentType   *GetOrderPaymentResponse `json:"payment_type"`
}

type GetOrderPaymentResponse struct {
	ID   int64  `json:"paymentTypeId"`
	Name string `json:"name"`
	Type string `json:"type"`
	Logo string `json:"logo"`
}

type CheckOrderDownloadResponse struct {
	IsDownload bool `json:"isDownload"`
}
