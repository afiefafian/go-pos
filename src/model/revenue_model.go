package model

type GetRevenueReportResponse struct {
	TotalRevenue int64  `json:"totalRevenue"`
	PaymentTypes string `json:"paymentTypes"`
}

type PaymentReportResponse struct {
	ID          int64  `json:"paymentTypeId"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Logo        string `json:"logo"`
	TotalAmount int64  `json:"totalAmount"`
}
