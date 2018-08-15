package origanizationservice

import (
	"github.com/scottxxx666/meetups-api/app"
	"github.com/scottxxx666/meetups-api/model"
)

func Find(id uint64) model.Origanization {
	var origanization model.Origanization
	if app.DB.Preload("Meetups").First(&origanization, id).RecordNotFound() {
		panic("Find no this origanization")
	}

	return origanization
}

func GetMeetups(id uint64) []model.Meetup {
	var origanization model.Origanization
	if app.DB.Preload("Meetups").First(&origanization, id).RecordNotFound() {
		panic("getMeetups no this origanization")
	}

	return origanization.Meetups
}
