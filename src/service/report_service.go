package service

import (
	"github.com/afiefafian/go-pos/src/model"
	"github.com/afiefafian/go-pos/src/repository"
)

type ReportService struct {
	PaymentRepository *repository.PaymentRepository
}

func NewReportService(paymentRepository *repository.PaymentRepository) *ReportService {
	return &ReportService{PaymentRepository: paymentRepository}
}

func (s *ReportService) GetPaymentTypeRevenues() (*model.GetRevenueReportResponse, error) {
	// Get data
	revenues, err := s.PaymentRepository.FindAllWithRevenues()
	if err != nil {
		return nil, err
	}

	var total int64 = 0
	for _, paymentType := range revenues {
		total += paymentType.TotalAmount
	}

	revenueResponse := &model.GetRevenueReportResponse{
		TotalRevenue: total,
		PaymentTypes: revenues,
	}
	return revenueResponse, nil
}

func (s *ReportService) GetOrderedProducts() (*model.GetOrderedProductReportResponse, error) {
	return nil, nil
}
