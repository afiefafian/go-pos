package entity

import (
	"time"
)

type OrderProduct struct {
	ID               int64
	OrderID          int64
	ProductID        int64
	DiscountID       int64
	Qty              int32
	Price            int64
	totalNormalPrice int64
	TotalFinalPrice  int64
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
