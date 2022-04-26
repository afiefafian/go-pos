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

func (r *CashierService) FindAll(params *model.PaginationQuery) ([]model.GetCashierResponse, *model.PaginationResponse, error) {
	// Get cashiers data
	cashiers, err := r.CashierRepository.FindAll(params)
	if err != nil {
		return nil, nil, err
	}

	// Get total cashiers
	var count int
	count, err = r.CashierRepository.Count()
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

func (r *CashierService) GetByID(id int64) (*model.GetCashierResponse, error) {
	cashier, err := r.CashierRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return &model.GetCashierResponse{
		ID:   cashier.ID,
		Name: cashier.Name,
	}, nil
}

func (r *CashierService) Create(request model.CreateCashierRequest) (*model.CreateCashierResponse, error) {
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

	id, err := r.CashierRepository.Create(cashier)
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

func (r *CashierService) UpdateByID(request model.UpdateCashierRequest) error {
	if _, err := r.CashierRepository.GetByID(request.ID); err != nil {
		return err
	}

	if err := helper.ValidateStruct(request); err != nil {
		return err
	}

	cashier := entity.Cashier{
		ID:       request.ID,
		Name:     request.Name,
		Passcode: request.Passcode,
	}

	if err := r.CashierRepository.UpdateByID(cashier); err != nil {
		return err
	}

	return nil
}

func (r *CashierService) DeleteByID(id int64) error {
	if _, err := r.CashierRepository.GetByID(id); err != nil {
		return err
	}

	if err := r.CashierRepository.DeleteByID(id); err != nil {
		return err
	}
	return nil
}
