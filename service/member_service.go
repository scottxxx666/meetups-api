package service

import (
	"github.com/scottxxx666/meetups-api/app"
	"github.com/scottxxx666/meetups-api/model"
)

// MemberService is the service of member
type MemberService struct{}

// Find find the mamber by ID
func (m *MemberService) Find(id uint64) model.Member {
	var member model.Member
	if app.DB.First(&member, id).RecordNotFound() {
		panic("no this member")
	}

	return member
}
