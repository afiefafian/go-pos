package entity

import "time"

type ProductDiscount struct {
	ID        int64
	Type      string
	Qty       int32
	Result    int64
	ExpiredAt int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
