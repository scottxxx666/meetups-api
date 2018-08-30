package resolver

import (
	"strconv"
	"time"

	"github.com/scottxxx666/meetups-api/app"
	"github.com/scottxxx666/meetups-api/model"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/scottxxx666/meetups-api/service/meetupservice"
)

type meetup struct {
	model.Meetup
}

// MeetupResolver resolve Meetup
type MeetupResolver struct {
	m *meetup
}

// ID resolve ID
func (r *MeetupResolver) ID() graphql.ID {
	id := strconv.FormatUint(r.m.ID, 10)
	return graphql.ID(id)
}

// CreatedAt resolve CreatedAt
func (r *MeetupResolver) CreatedAt() string {
	return r.m.CreatedAt.String()
}

// UpdatedAt resolve UpdatedAt
func (r *MeetupResolver) UpdatedAt() string {
	return r.m.UpdatedAt.String()
}

// Name resolve Name
func (r *MeetupResolver) Name() string {
	return r.m.Name
}

// StartTime resolve StartTime
func (r *MeetupResolver) StartTime() string {
	return r.m.StartTime.Format(app.Timeformat)
}

// EndTime resolve EndTime
func (r *MeetupResolver) EndTime() string {
	return r.m.EndTime.Format(app.Timeformat)
}

// Rating resolve Rating
func (r *MeetupResolver) Rating() float64 {
	return r.m.Rating
}

// RatingCount resolve RatingCount
func (r *MeetupResolver) RatingCount() int32 {
	return r.m.RatingCount
}

// NormalPrice resolve NormalPrice
func (r *MeetupResolver) NormalPrice() int32 {
	return r.m.NormalPrice
}

// Organization resolve Organization
func (r *MeetupResolver) Organization() *OrganizationResolver {
	return &OrganizationResolver{&organization{r.m.Organization}}
}

// Level resolve Level
func (r *MeetupResolver) Level() string {
	return r.m.Level.Name
}

// Location resolve Location
func (r *MeetupResolver) Location() string {
	return r.m.Location.Name
}

// Tags resolve Tags
func (r *MeetupResolver) Tags() []string {
	var t []string
	for _, tag := range r.m.Tags {
		t = append(t, tag.ID)
	}
	return t
}

// Reviews resolve Reviews
func (r *MeetupResolver) Reviews() []*ReviewResolver {
	var result []*ReviewResolver
	for _, rev := range r.m.Reviews {
		result = append(result, &ReviewResolver{&review{rev}})
	}
	return result
}

// Meetup resolve query Meetup
func (r *Resolver) Meetup(args struct{ ID string }) *MeetupResolver {
	id, err := strconv.ParseUint(args.ID, 10, 64)
	if err != nil {
		return nil
	}
	result := meetupservice.Find(id)
	m := meetup{result}
	return &MeetupResolver{&m}
}

type meetupArgs struct {
	Meetup meetupInput
}

type meetupInput struct {
	Name           string
	StartTime      string
	EndTime        string
	NormalPrice    float64
	OrganizationID float64
	LevelID        float64
	LocationID     float64
	Tags           []string
}

// CreateMeetup resolve query createMeetup
// func (r *Resolver) CreateMeetup(args meetupservice.MeetupInput) *MeetupResolver {
func (r *Resolver) CreateMeetup(args meetupArgs) *MeetupResolver {
	startTime := parseTime(args.Meetup.StartTime)
	endTime := parseTime(args.Meetup.EndTime)
	m := meetupservice.Create(meetupservice.MeetupArgs{
		Name:           args.Meetup.Name,
		StartTime:      startTime,
		EndTime:        endTime,
		NormalPrice:    int32(args.Meetup.NormalPrice),
		OrganizationID: uint64(args.Meetup.OrganizationID),
		LevelID:        uint64(args.Meetup.LevelID),
		LocationID:     uint64(args.Meetup.LocationID),
		Tags:           args.Meetup.Tags,
	})
	return &MeetupResolver{&meetup{m}}
}

func parseTime(s string) time.Time {
	t, err := time.Parse(app.Timeformat, "2018-12-25 00:00")
	if err != nil {
		panic(err)
	}
	return t
}
