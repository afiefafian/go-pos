package entity

import "time"

type Cashier struct {
	ID        int
	Name      string
	Passcode  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
