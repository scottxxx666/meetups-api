package reviewservice

import (
	"github.com/scottxxx666/meetups-api/app"
	"github.com/scottxxx666/meetups-api/model"
)

func Find(id uint64) model.Review {
	var review model.Review
	if app.DB.Preload("Member").First(&review, id).RecordNotFound() {
		panic("Find no this origanization")
	}

	return review
}

func GetByMeetup(id uint64) []model.Review {
	var reviews []model.Review
	if app.DB.Preload("Member").Where("meetup_id = ?", id).Find(&reviews).RecordNotFound() {
		panic("Find by meetup id failed")
	}
	return reviews
}
