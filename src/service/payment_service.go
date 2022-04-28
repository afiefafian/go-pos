package service

import (
	"math"
	"time"

	"github.com/afiefafian/go-pos/src/entity"
	"github.com/afiefafian/go-pos/src/helper"
	"github.com/afiefafian/go-pos/src/model"
	"github.com/afiefafian/go-pos/src/repository"
)

type PaymentService struct {
	PaymentRepository *repository.PaymentRepository
}

func NewPaymentService(paymentRepository *repository.PaymentRepository) *PaymentService {
	return &PaymentService{PaymentRepository: paymentRepository}
}

func (s *PaymentService) FindAll(params model.GetPaymentQuery) ([]model.GetPaymentResponse, *model.PaginationResponse, error) {
	// Get data
	payments, err := s.PaymentRepository.FindAll(&params.Pagination)
	if err != nil {
		return nil, nil, err
	}

	// Get total
	var count int
	count, err = s.PaymentRepository.Count()
	if err != nil {
		return nil, nil, err
	}

	// Format data
	var paymentsResponse = make([]model.GetPaymentResponse, 0)
	for _, row := range payments {
		payment := model.GetPaymentResponse{
			ID:   row.ID,
			Name: row.Name,
			Type: row.Type,
			Logo: row.Logo,
		}

		if payment.Type == "CASH" && params.Subtotal > 0 {
			payment.Card = append(payment.Card, params.Subtotal)

			// Add +1 if modulus result is 0
			subtotalSuggestion := params.Subtotal
			if subtotalSuggestion%10000 == 0 {
				subtotalSuggestion = subtotalSuggestion + 1
			}

			// Round subtotal to nearest 10k
			roundedSubtotal := math.Ceil(float64(subtotalSuggestion)/10000) * 10000
			payment.Card = append(payment.Card, int64(roundedSubtotal))
		}

		paymentsResponse = append(paymentsResponse, payment)
	}

	pagination := &model.PaginationResponse{
		Total: count,
		Skip:  params.Pagination.Skip,
		Limit: params.Pagination.Limit,
	}

	return paymentsResponse, pagination, nil
}

func (s *PaymentService) GetByID(id int64) (*model.GetPaymentResponse, error) {
	payment, err := s.PaymentRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return &model.GetPaymentResponse{
		ID:   payment.ID,
		Name: payment.Name,
		Type: payment.Type,
		Logo: payment.Logo,
	}, nil
}

func (s *PaymentService) Create(request model.CreatePaymentRequest) (*model.CreatePaymentResponse, error) {
	if err := helper.ValidateStruct(request); err != nil {
		return nil, err
	}

	currentTime := time.Now()
	payment := entity.Payment{
		Name:      request.Name,
		Type:      request.Type,
		Logo:      request.Logo,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}

	id, err := s.PaymentRepository.Create(payment)
	if err != nil {
		return nil, err
	}

	return &model.CreatePaymentResponse{
		ID:        id,
		Name:      payment.Name,
		Type:      request.Type,
		Logo:      request.Logo,
		CreatedAt: payment.CreatedAt,
		UpdatedAt: payment.UpdatedAt,
	}, nil
}

func (s *PaymentService) UpdateByID(request model.UpdatePaymentRequest) error {
	existingPayment, err := s.PaymentRepository.GetByID(request.ID)
	if err != nil {
		return err
	}

	if err := helper.ValidateStruct(request); err != nil {
		return err
	}

	payment := entity.Payment{
		ID:   request.ID,
		Name: request.Name,
		Type: request.Type,
		Logo: request.Logo,
	}

	// Set default value
	if request.Name == "" {
		payment.Name = existingPayment.Name
	}
	if request.Type == "" {
		payment.Type = existingPayment.Type
	}
	if request.Logo == "" {
		payment.Logo = existingPayment.Logo
	}

	if err := s.PaymentRepository.UpdateByID(payment); err != nil {
		return err
	}

	return nil
}

func (s *PaymentService) DeleteByID(id int64) error {
	if _, err := s.PaymentRepository.GetByID(id); err != nil {
		return err
	}

	if err := s.PaymentRepository.DeleteByID(id); err != nil {
		return err
	}
	return nil
}
