package resolver

import (
	"github.com/scottxxx666/meetups-api/service/meetupservice"
)

// HotMeetups resolve query hotMeetups
func (r *Resolver) HotMeetups() []*MeetupResolver {
	var mrs []*MeetupResolver
	meetups := meetupservice.GetHotMeetups()
	for _, m := range meetups {
		mrs = append(mrs, &MeetupResolver{&meetup{m}})
	}
	return mrs
}

// Meetups resolve query Meetups
func (r *Resolver) Meetups() []*MeetupResolver {
	var mrs []*MeetupResolver
	meetups := meetupservice.Get()
	for _, m := range meetups {
		mrs = append(mrs, &MeetupResolver{&meetup{m}})
	}
	return mrs
}
