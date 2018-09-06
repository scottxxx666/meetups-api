package service

import (
	"time"

	"github.com/scottxxx666/meetups-api/app"
	"github.com/scottxxx666/meetups-api/model"
)

// MeetupService is the service of meetup
type MeetupService struct{}

// Find meetup by ID
func (m *MeetupService) Find(id uint64) model.Meetup {
	var meetup model.Meetup
	if app.DB.Preload("Level").Preload("Location").Preload("Tags").Preload("Organization").First(&meetup, id).RecordNotFound() {
		panic("no this meetup")
	}

	return meetup
}

// GetHotMeetups get the hot meetups
func (m *MeetupService) GetHotMeetups() []model.Meetup {
	var meetups []model.Meetup
	if app.DB.Preload("Level").Preload("Location").Preload("Tags").Preload("Organization").Limit(3).Find(&meetups).RecordNotFound() {
		panic("no this meetup")
	}

	return meetups
}

// Get get the meetups
func (m *MeetupService) Get() []model.Meetup {
	var meetups []model.Meetup
	if app.DB.Preload("Level").Preload("Location").Preload("Tags").Preload("Organization").Limit(3).Find(&meetups).RecordNotFound() {
		panic("no this meetup")
	}

	return meetups
}

// MeetupArgs is the input for meetup create
type MeetupArgs struct {
	Name           string
	StartTime      time.Time
	EndTime        time.Time
	NormalPrice    int32
	OrganizationID uint64
	LevelID        uint64
	LocationID     uint64
	Tags           []string
}

// Create create a meetup
func (m *MeetupService) Create(args MeetupArgs) model.Meetup {
	var ts []model.Tag

	for _, t := range args.Tags {
		ts = append(ts, model.Tag{ID: t})
	}

	meetup := model.Meetup{
		Name:           args.Name,
		StartTime:      args.StartTime,
		EndTime:        args.EndTime,
		NormalPrice:    args.NormalPrice,
		OrganizationID: args.OrganizationID,
		LevelID:        args.LevelID,
		LocationID:     args.LocationID,
		Tags:           ts,
	}

	if app.DB.Create(&meetup).Error != nil {
		panic("meetup create failed")
	}

	return meetup
}
