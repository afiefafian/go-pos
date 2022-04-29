package service

import (
	"github.com/afiefafian/go-pos/src/model"
	"github.com/afiefafian/go-pos/src/repository"
)

type ReportService struct {
	PaymentRepository *repository.PaymentRepository
	ProductRepository *repository.ProductRepository
}

func NewReportService(
	paymentRepository *repository.PaymentRepository,
	productRepository *repository.ProductRepository,
) *ReportService {
	return &ReportService{
		PaymentRepository: paymentRepository,
		ProductRepository: productRepository,
	}
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

	response := &model.GetRevenueReportResponse{
		TotalRevenue: total,
		PaymentTypes: revenues,
	}
	return response, nil
}

func (s *ReportService) GetOrderedProducts() ([]model.GetOrderedProductResponse, error) {
	products, err := s.ProductRepository.FindAllOrderedProducts()
	if err != nil {
		return nil, err
	}

	return products, nil
}
