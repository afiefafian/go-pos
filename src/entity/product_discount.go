package entity

import (
	"fmt"
	"time"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
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
	if d.ExpiredAt != 0 {
		date := time.Unix(d.ExpiredAt, 0).UTC()
		return &date
	}
	return nil
}

func (d *ProductDiscount) ExpiredAtFormat() string {
	if d.ExpiredAt != 0 {
		date := time.Unix(d.ExpiredAt, 0).UTC()
		return date.Format("01 Jan 2006")
	}
	return ""
}

func (d *ProductDiscount) StringFormat() string {
	switch d.Type {
	case "BUY_N":
		return fmt.Sprintf("Buy %d only Rp. %d", d.Qty, d.Result)
	case "PERCENT":
		p := message.NewPrinter(language.Indonesian)
		formatResult := p.Sprintf("%d", 1000)
		return fmt.Sprintf("Discount %d%% only Rp. %s", d.Qty, formatResult)
	default:
		return ""
	}

}

func (d *ProductDiscount) IsValid() bool {
	return true
}
