package model

import "time"

type Origanization struct {
	ID        uint64
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Meetups   []Meetup
}
