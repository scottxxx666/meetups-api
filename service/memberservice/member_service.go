package memberservice

import (
	"github.com/scottxxx666/meetups-api/app"
	"github.com/scottxxx666/meetups-api/model"
)

func Find(id uint64) model.Member {
	var member model.Member
	if app.DB.First(&member, id).RecordNotFound() {
		panic("no this member")
	}

	return member
}
