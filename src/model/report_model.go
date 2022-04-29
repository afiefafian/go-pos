package model

type GetRevenueReportResponse struct {
	TotalRevenue int64                   `json:"totalRevenue"`
	PaymentTypes []PaymentReportResponse `json:"paymentTypes"`
}

type PaymentReportResponse struct {
	ID          int64  `json:"paymentTypeId"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Logo        string `json:"logo"`
	TotalAmount int64  `json:"totalAmount"`
}

type GetOrderedProductResponse struct {
	ID          int64  `json:"productId"`
	Name        string `json:"name"`
	TotalQty    int64  `json:"totalQty"`
	TotalAmount int64  `json:"totalAmount"`
}
