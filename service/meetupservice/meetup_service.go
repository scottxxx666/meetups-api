package meetupservice

import (
	"github.com/scottxxx666/meetups-api/app"
	"github.com/scottxxx666/meetups-api/model"
)

func Find(id uint64) model.Meetup {
	var meetup model.Meetup
	if app.DB.Preload("Level").Preload("Location").Preload("Tags").First(&meetup, id).RecordNotFound() {
		panic("no this member")
	}

	return meetup
}
