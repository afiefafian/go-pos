package model

import "time"

type GetOrderResponse struct {
	ID            int64                    `json:"orderId"`
	CashierID     int64                    `json:"cashiersId,omitempty"`
	PaymentTypeID int64                    `json:"paymentTypesId,omitempty"`
	TotalPrice    int64                    `json:"totalPrice"`
	TotalPaid     int64                    `json:"totalPaid"`
	TotalReturn   int64                    `json:"totalReturn"`
	ReceiptID     string                   `json:"receiptId"`
	CreatedAt     time.Time                `json:"createdAt"`
	UpdatedAt     time.Time                `json:"updatedAt,omitempty"`
	Cashier       *GetCashierResponse      `json:"cashier,omitempty"`
	PaymentType   *GetOrderPaymentResponse `json:"payment_type,omitempty"`
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

type GetDetailOrderResponse CreateOrderResponse
