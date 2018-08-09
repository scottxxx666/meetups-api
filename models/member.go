package models

import "time"

type Member struct {
	ID        uint64
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
}
