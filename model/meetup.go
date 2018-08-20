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
	OriganizationID uint64
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Level           Level
	LevelID         uint64
	Location        Location
	LocationID      uint64
	Tags            []Tag
	Reviews         []Review
}
