package entity

import (
	"fmt"
	"math/rand"
	"time"
)

type Order struct {
	ID            int64
	CashierID     int64
	PaymentTypeID int64
	TotalPrice    int64
	TotalPaid     int64
	TotalReturn   int64
	ReceiptID     string
	IsDownloaded  bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (d *Order) GenerateReceiptID() string {
	// Generate random number
	hi := 999
	low := 100
	randNum := low + rand.Intn(hi-low)

	// Generate random string
	randStr := "A"

	return fmt.Sprintf("S%d%s", randNum, randStr)
}
