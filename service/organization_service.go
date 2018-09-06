package service

import (
	"github.com/scottxxx666/meetups-api/app"
	"github.com/scottxxx666/meetups-api/model"
)

// OrganizationService is the service of organization
type OrganizationService struct{}

func (o *OrganizationService) All() []model.Organization {
	var organizations []model.Organization
	if app.DB.Find(&organizations).Error != nil {
		panic("Get all organization failed")
	}
	return organizations
}

// Find find the organization by ID
func (o *OrganizationService) Find(id uint64) model.Organization {
	var organization model.Organization
	if app.DB.First(&organization, id).RecordNotFound() {
		panic("Find no this organization")
	}

	return organization
}

// GetMeetups get the organization by ID with its meetups
func (o *OrganizationService) GetMeetups(id uint64) []model.Meetup {
	var organization model.Organization
	if app.DB.Preload("Meetups").First(&organization, id).RecordNotFound() {
		panic("getMeetups no this organization")
	}

	return organization.Meetups
}
