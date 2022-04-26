package service

import (
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

func (s *PaymentService) FindAll(params *model.PaginationQuery) ([]model.GetPaymentResponse, *model.PaginationResponse, error) {
	// Get data
	payments, err := s.PaymentRepository.FindAll(params)
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
	for _, payment := range payments {
		paymentsResponse = append(paymentsResponse, model.GetPaymentResponse{
			ID:   payment.ID,
			Name: payment.Name,
			Type: payment.Type,
			Logo: payment.Logo,
		})
	}

	pagination := &model.PaginationResponse{
		Total: count,
		Skip:  params.Skip,
		Limit: params.Limit,
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
	if _, err := s.PaymentRepository.GetByID(request.ID); err != nil {
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
