package service

import (
	"time"

	"github.com/afiefafian/go-pos/src/entity"
	"github.com/afiefafian/go-pos/src/helper"
	"github.com/afiefafian/go-pos/src/model"
	"github.com/afiefafian/go-pos/src/repository"
)

type CashierService struct {
	CashierRepository *repository.CashierRepository
}

func NewCashierService(cashierRepository *repository.CashierRepository) *CashierService {
	return &CashierService{CashierRepository: cashierRepository}
}

func (s *CashierService) FindAll(params *model.PaginationQuery) ([]model.GetCashierResponse, *model.PaginationResponse, error) {
	// Get cashiers data
	cashiers, err := s.CashierRepository.FindAll(params)
	if err != nil {
		return nil, nil, err
	}

	// Get total cashiers
	var count int
	count, err = s.CashierRepository.Count()
	if err != nil {
		return nil, nil, err
	}

	// Format cashiers data
	var cashiersResponse = make([]model.GetCashierResponse, 0)
	for _, cashier := range cashiers {
		cashiersResponse = append(cashiersResponse, model.GetCashierResponse{
			ID:   cashier.ID,
			Name: cashier.Name,
		})
	}

	pagination := &model.PaginationResponse{
		Total: count,
		Skip:  params.Skip,
		Limit: params.Limit,
	}

	return cashiersResponse, pagination, nil
}

func (s *CashierService) GetByID(id int64) (*model.GetCashierResponse, error) {
	cashier, err := s.CashierRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return &model.GetCashierResponse{
		ID:   cashier.ID,
		Name: cashier.Name,
	}, nil
}

func (s *CashierService) Create(request model.CreateCashierRequest) (*model.CreateCashierResponse, error) {
	if err := helper.ValidateStruct(request); err != nil {
		return nil, err
	}

	currentTime := time.Now()
	cashier := entity.Cashier{
		Name:      request.Name,
		Passcode:  request.Passcode,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}

	id, err := s.CashierRepository.Create(cashier)
	if err != nil {
		return nil, err
	}

	return &model.CreateCashierResponse{
		ID:        id,
		Name:      cashier.Name,
		Passcode:  cashier.Passcode,
		CreatedAt: cashier.CreatedAt,
		UpdatedAt: cashier.UpdatedAt,
	}, nil
}

func (s *CashierService) UpdateByID(request model.UpdateCashierRequest) error {
	if _, err := s.CashierRepository.GetByID(request.ID); err != nil {
		return err
	}

	if err := helper.ValidateStruct(request); err != nil {
		return err
	}

	cashier := entity.Cashier{
		ID:   request.ID,
		Name: request.Name,
	}

	if err := s.CashierRepository.UpdateByID(cashier); err != nil {
		return err
	}

	return nil
}

func (s *CashierService) DeleteByID(id int64) error {
	if _, err := s.CashierRepository.GetByID(id); err != nil {
		return err
	}

	if err := s.CashierRepository.DeleteByID(id); err != nil {
		return err
	}
	return nil
}
