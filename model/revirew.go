package model

import "time"

type Review struct {
	ID        uint64
	Rating    int32
	Desc      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Member    Member
	MemberID  uint64
	Meetup    Meetup
	MeetupID  uint64
}
