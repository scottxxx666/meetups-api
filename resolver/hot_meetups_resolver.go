package resolver

import (
	"github.com/scottxxx666/meetups-api/service/meetupservice"
)

func (r *Resolver) HotMeetups() []*MeetupResolver {
	var mrs []*MeetupResolver
	meetups := meetupservice.GetHotMeetups()
	for _, m := range meetups {
		mrs = append(mrs, &MeetupResolver{&meetup{m}})
	}
	return mrs
}
