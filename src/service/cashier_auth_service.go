package service

import (
	"strconv"

	"github.com/afiefafian/go-pos/src/exception"
	"github.com/afiefafian/go-pos/src/helper"
	"github.com/afiefafian/go-pos/src/model"
	"github.com/afiefafian/go-pos/src/repository"
)

type CashierAuthService struct {
	CashierRepository *repository.CashierRepository
}

func NewCashierAuthService(cashierAuthRepository *repository.CashierRepository) *CashierAuthService {
	return &CashierAuthService{
		CashierRepository: cashierAuthRepository,
	}
}

func (s *CashierAuthService) GetPasscode(id int64) (*model.CashierPasscodeRespose, error) {
	cashier, err := s.CashierRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return &model.CashierPasscodeRespose{
		ID:       id,
		Passcode: cashier.Passcode,
	}, nil
}

func (s *CashierAuthService) Authenticate(request model.CashierLoginRequest) (*model.CashierLoginResponse, error) {
	cashier, err := s.CashierRepository.GetByID(request.ID)
	if err != nil {
		return nil, err
	}

	if err := helper.ValidateStruct(request); err != nil {
		return nil, err
	}

	if cashier.Passcode != request.Passcode {
		return nil, exception.Unauthorized(helper.MsgPasscodeNotMatch)
	}

	// Generate token
	var token string
	strId := strconv.FormatInt(cashier.ID, 10)
	token, err = helper.GenerateJWToken(strId)
	if err != nil {
		return nil, err
	}

	return &model.CashierLoginResponse{
		Token: token,
	}, nil
}

func (s *CashierAuthService) Logout(request model.CashierLogoutRequest) error {
	cashier, err := s.CashierRepository.GetByID(request.ID)
	if err != nil {
		return err
	}

	if err := helper.ValidateStruct(request); err != nil {
		return err
	}

	if cashier.Passcode != request.Passcode {
		return exception.Unauthorized(helper.MsgPasscodeNotMatch)
	}

	return nil
}
