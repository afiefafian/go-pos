package service

import (
	"github.com/afiefafian/go-pos/src/repository"
)

type CashierService struct {
	cashierRepository *repository.CashierRepository
}

func NewCashierService(cashierRepository *repository.CashierRepository) *CashierService {
	return &CashierService{cashierRepository}
}
