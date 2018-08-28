package meetupservice

import (
	"github.com/scottxxx666/meetups-api/app"
	"github.com/scottxxx666/meetups-api/model"
)

func Find(id uint64) model.Meetup {
	var meetup model.Meetup
	if app.DB.Preload("Level").Preload("Location").Preload("Tags").Preload("Organization").First(&meetup, id).RecordNotFound() {
		panic("no this meetup")
	}

	return meetup
}

func GetHotMeetups() []model.Meetup {
	var meetups []model.Meetup
	if app.DB.Preload("Level").Preload("Location").Preload("Tags").Preload("Organization").Limit(3).Find(&meetups).RecordNotFound() {
		panic("no this meetup")
	}

	return meetups
}
