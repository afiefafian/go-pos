package entity

import (
	"time"
)

type OrderProductSnapshot struct {
	ID         int64
	OrderID    int64
	ProductID  int64
	CategoryID int64
	DiscountID int64
	Name       string
	Stock      int32
	Price      int64
	Image      string
	Category   *string
	Discount   *string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
