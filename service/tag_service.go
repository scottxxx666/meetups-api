package service

import (
	"github.com/scottxxx666/meetups-api/app"
	"github.com/scottxxx666/meetups-api/model"
)

// TagService is a service of tag
type TagService struct{}

// All get all tags
func (t *TagService) All() []model.Tag {
	var tags []model.Tag
	if app.DB.Find(&tags).Error != nil {
		panic("Get all tags error")
	}
	return tags
}
