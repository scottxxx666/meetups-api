package model

import "time"

type Origanization struct {
	ID        uint64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Meetups   []Meetup
}
