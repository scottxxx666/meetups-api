package model

import "time"

type Origanization struct {
	ID        uint64
	Name      string
	Meetups   []Meetup
	CreatedAt time.Time
	UpdatedAt time.Time
}
