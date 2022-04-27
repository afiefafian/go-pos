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
