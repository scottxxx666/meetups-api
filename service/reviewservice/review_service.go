package reviewservice

import (
	"github.com/scottxxx666/meetups-api/app"
	"github.com/scottxxx666/meetups-api/model"
)

func Find(id uint64) model.Review {
	var review model.Review
	if app.DB.Preload("Member").First(&review, id).RecordNotFound() {
		panic("Find no this organization")
	}

	return review
}

func GetByMeetup(id uint64, first int) ([]model.Review, int32) {
	var reviews []model.Review
	var count int32
	if app.DB.Preload("Member").Where("meetup_id = ?", id).Order("updated_at desc, id desc").Limit(first).Find(&reviews).Count(&count).RecordNotFound() {
		panic("Find by meetup id failed")
	}
	return reviews, count
}

func GetByMeetupAfter(id uint64, first int, afterID uint64, afterUpdatedAt string) ([]model.Review, int32) {
	var reviews []model.Review
	var count int32
	if app.DB.Preload("Member").Where("meetup_id = ?", id).Where("updated_at <= ?", afterUpdatedAt).Not("updated_at = ? ANd id >= ?", afterUpdatedAt, afterID).Order("updated_at desc, id desc").Limit(first).Find(&reviews).Count(&count).RecordNotFound() {
		panic("Find by meetup id failed")
	}
	return reviews, count
}
