package model

type CreateCashierRequest struct {
	Name     string `json:"name"`
	Passcode string `json:"passcode"`
}

type GetCashierResponse CreateCashierRequest

type CreateCashierResponse struct {
	ID        string `json:"cashierId"`
	Name      string `json:"name"`
	Passcode  string `json:"passcode"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type CashierLoginRequest struct {
	Passcode string `json:"passcode"`
}

type CashierLogoutRequest CashierLoginRequest

type CashierLoginResponse struct {
	Token string `json:"token"`
}
