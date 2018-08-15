package model

import "time"

type Tag struct {
	ID        uint64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	MeetupID  uint64
	Meetup    Meetup
}
