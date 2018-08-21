package model

import "time"

type Meetup struct {
	ID              uint64
	Name            string
	StartTime       time.Time
	EndTime         time.Time
	Rating          float64
	RatingCount     int32
	NormalPrice     int32
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Origanization   Origanization
	OriganizationID uint64
	Level           Level
	LevelID         uint64
	Location        Location
	LocationID      uint64
	Tags            []Tag
	Reviews         []Review
}
