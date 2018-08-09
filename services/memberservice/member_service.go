package memberservice

import (
	"github.com/scottxxx666/meetups-api/app"

	"github.com/scottxxx666/meetups-api/models"
)

func Find(id uint64) models.Member {
	var member models.Member
	if app.DB.First(&member, id).RecordNotFound() {
		panic("no this member")
	}

	return member
}
