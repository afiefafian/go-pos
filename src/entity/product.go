package entity

import "time"

type Product struct {
	ID         int64
	Name       string
	Stock      int32
	Price      int
	Image      string
	CategoryID int64
	DiscountID int64
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
