package service

import (
	"github.com/scottxxx666/meetups-api/app"
	"github.com/scottxxx666/meetups-api/model"
)

// LocationService is the service of location
type LocationService struct{}

// All get all locations
func (s *LocationService) All() []model.Location {
	var locations []model.Location
	if app.DB.Find(&locations).Error != nil {
		panic("Get all locations error")
	}
	return locations
}
