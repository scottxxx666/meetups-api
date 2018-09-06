package resolver

import (
	"github.com/scottxxx666/meetups-api/service"
)

// HotMeetups resolve query hotMeetups
func (r *Resolver) HotMeetups() []*MeetupResolver {
	var mrs []*MeetupResolver
	var ms service.MeetupService
	meetups := ms.GetHotMeetups()
	for _, m := range meetups {
		mrs = append(mrs, &MeetupResolver{&meetup{m}})
	}
	return mrs
}

// Meetups resolve query Meetups
func (r *Resolver) Meetups() []*MeetupResolver {
	var mrs []*MeetupResolver
	var ms service.MeetupService
	meetups := ms.Get()
	for _, m := range meetups {
		mrs = append(mrs, &MeetupResolver{&meetup{m}})
	}
	return mrs
}
