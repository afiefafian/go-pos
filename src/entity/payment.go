package entity

import "time"

type Payment struct {
	ID        int64
	Name      string
	Type      string
	Logo      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
