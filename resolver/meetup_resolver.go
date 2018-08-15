package resolver

import (
	"strconv"

	"github.com/scottxxx666/meetups-api/model"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/scottxxx666/meetups-api/service/meetupservice"
)

type meetup struct {
	model.Meetup
}

type MeetupResolver struct {
	m *meetup
}

func (r *MeetupResolver) ID() graphql.ID {
	id := strconv.FormatUint(r.m.ID, 10)
	return graphql.ID(id)
}

func (r *MeetupResolver) CreatedAt() string {
	return r.m.CreatedAt.String()
}

func (r *MeetupResolver) UpdatedAt() string {
	return r.m.UpdatedAt.String()
}

func (r *MeetupResolver) Name() string {
	return r.m.Name
}

func (r *MeetupResolver) StartTime() string {
	return r.m.StartTime.String()
}

func (r *MeetupResolver) EndTime() string {
	return r.m.EndTime.String()
}

func (r *MeetupResolver) Rating() float64 {
	return r.m.Rating
}

func (r *MeetupResolver) RatingCount() int32 {
	return r.m.RatingCount
}

func (r *MeetupResolver) NormalPrice() int32 {
	return r.m.NormalPrice
}

func (r *Resolver) Meetup(args struct{ ID string }) *MeetupResolver {
	id, err := strconv.ParseUint(args.ID, 10, 64)
	if err != nil {
		return nil
	}
	result := meetupservice.Find(id)
	m := meetup{result}
	return &MeetupResolver{&m}
}
