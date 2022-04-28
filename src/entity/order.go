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

func (o *Order) GenerateReceiptID() string {
	// Generate random number
	low := 0
	hi := 999
	randNum := low + rand.Intn(hi-low)

	// Generate random string
	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	randStr := string(charset[rand.Intn(len(charset))])

	return fmt.Sprintf("S%03d%s", randNum, randStr)
}
