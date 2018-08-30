package model

import (
	"time"
)

type Tag struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	Meetups   []Meetup `gorm:"many2many:meetup_tags;"`
}
