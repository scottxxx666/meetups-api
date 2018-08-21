package organizationservice

import (
	"github.com/scottxxx666/meetups-api/app"
	"github.com/scottxxx666/meetups-api/model"
)

func Find(id uint64) model.Organization {
	var organization model.Organization
	if app.DB.First(&organization, id).RecordNotFound() {
		panic("Find no this organization")
	}

	return organization
}

func GetMeetups(id uint64) []model.Meetup {
	var organization model.Organization
	if app.DB.Preload("Meetups").First(&organization, id).RecordNotFound() {
		panic("getMeetups no this organization")
	}

	return organization.Meetups
}
