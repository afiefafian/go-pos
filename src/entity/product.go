package entity

import (
	"fmt"
	"time"
)

type Product struct {
	ID        int64
	Name      string
	Stock     int32
	Price     int64
	Image     string
	Category  *Category
	Discount  *ProductDiscount
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p *Product) SKU() string {
	return fmt.Sprintf("ID%03d", p.ID)
}

func (p *Product) FinalPrice() int64 {
	if p.Discount != nil && p.Discount.IsValid() {
		switch p.Discount.Type {
		case "BUY_N":
			return p.Discount.Result
		case "PERCENT":
			return p.Price - ((p.Price * p.Discount.Result) / 100)
		default:
			return p.Price
		}
	}
	return p.Price
}
