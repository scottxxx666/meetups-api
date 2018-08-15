package model

import "time"

type Member struct {
	ID        uint64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
