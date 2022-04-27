package entity

import (
	"time"
)

type ProductDiscount struct {
	ID        int64
	Type      string
	Qty       int32
	Result    int64
	ExpiredAt int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (d *ProductDiscount) ExpiredAtDate() *time.Time {
	return nil
}

func (d *ProductDiscount) ExpiredAtFormat() string {
	return ""
}

func (d *ProductDiscount) StringFormat() string {
	return ""
	// return fmt.Sprintf("Buy %s only Rp. %s", d.Qty)
}
