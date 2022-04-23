package service

import (
	"github.com/afiefafian/go-pos/src/repository"
)

type CashierAuthService struct {
	cashierRepository *repository.CashierRepository
}

func NewCashierAuthService(cashierAuthRepository *repository.CashierRepository) *CashierAuthService {
	return &CashierAuthService{cashierAuthRepository}
}

func (s *CashierAuthService) GetCashierPasscode() {
}

func (s *CashierAuthService) Authenticate() {
}

func (s *CashierAuthService) Logout() {
}
