package meetupservice

import (
	"github.com/scottxxx666/meetups-api/app"
	"github.com/scottxxx666/meetups-api/model"
)

// Find meetup by ID
func Find(id uint64) model.Meetup {
	var meetup model.Meetup
	if app.DB.Preload("Level").Preload("Location").Preload("Tags").Preload("Organization").First(&meetup, id).RecordNotFound() {
		panic("no this meetup")
	}

	return meetup
}

// GetHotMeetups get the hot meetups
func GetHotMeetups() []model.Meetup {
	var meetups []model.Meetup
	if app.DB.Preload("Level").Preload("Location").Preload("Tags").Preload("Organization").Limit(3).Find(&meetups).RecordNotFound() {
		panic("no this meetup")
	}

	return meetups
}

// Get get the meetups
func Get() []model.Meetup {
	var meetups []model.Meetup
	if app.DB.Preload("Level").Preload("Location").Preload("Tags").Preload("Organization").Limit(3).Find(&meetups).RecordNotFound() {
		panic("no this meetup")
	}

	return meetups
}
