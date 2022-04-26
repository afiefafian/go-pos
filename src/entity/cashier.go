package entity

import "time"

type Cashier struct {
	ID        int64
	Name      string
	Passcode  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
