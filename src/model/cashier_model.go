package model

import "time"

type CreateCashierRequest struct {
	Name     string `json:"name" validate:"required,min=1,max=50"`
	Passcode string `json:"passcode" validate:"required,min=3,max=32"`
}

type UpdateCashierRequest struct {
	ID       int64
	Name     string `json:"name" validate:"required,min=1,max=50"`
	Passcode string `json:"passcode" validate:"required,min=3,max=32"`
}

type CashierLoginRequest struct {
	ID       int64
	Passcode string `json:"passcode" validate:"required,min=3,max=32"`
}

type CashierLogoutRequest CashierLoginRequest

type GetCashierResponse struct {
	ID   int64  `json:"cashierId"`
	Name string `json:"name"`
}

type CreateCashierResponse struct {
	ID        int64     `json:"cashierId"`
	Name      string    `json:"name"`
	Passcode  string    `json:"passcode"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type CashierPasscodeRespose CashierLoginRequest

type CashierLoginResponse struct {
	Token string `json:"token"`
}
