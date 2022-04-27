package service

import (
	"github.com/afiefafian/go-pos/src/helper"
	"github.com/afiefafian/go-pos/src/model"
	"github.com/afiefafian/go-pos/src/repository"
)

type OrderService struct {
	OrderRepository   *repository.OrderRepository
	ProductRepository *repository.ProductRepository
}

func NewOrderService(
	orderRepository *repository.OrderRepository,
	paymentRepository *repository.ProductRepository,
) *OrderService {
	return &OrderService{
		OrderRepository:   orderRepository,
		ProductRepository: paymentRepository,
	}
}

func (s *OrderService) CheckSubTotal(request model.CreateSubTotalRequest) (*model.CreateSubTotalResponse, error) {
	if err := helper.ValidateStruct(request); err != nil {
		return nil, err
	}

	return nil, nil
}
