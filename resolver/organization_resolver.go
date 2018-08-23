package resolver

import (
	"strconv"

	"github.com/scottxxx666/meetups-api/model"
	"github.com/scottxxx666/meetups-api/service/organizationservice"

	graphql "github.com/graph-gophers/graphql-go"
)

type organization struct {
	model.Organization
}

// OrganizationResolver resolve Organization
type OrganizationResolver struct {
	o *organization
}

// ID resolver
func (r *OrganizationResolver) ID() graphql.ID {
	id := strconv.FormatUint(r.o.ID, 10)
	return graphql.ID(id)
}

// Name resolver
func (r *OrganizationResolver) Name() string {
	return r.o.Name
}

// Meetups resolver
func (r *OrganizationResolver) Meetups() []*MeetupResolver {
	result := organizationservice.GetMeetups(r.o.ID)
	var mr []*MeetupResolver
	for _, m := range result {
		mr = append(mr, &MeetupResolver{&meetup{m}})
	}
	return mr
}

// Organization resolve query Organization
func (r *Resolver) Organization(args struct{ ID string }) *OrganizationResolver {
	id, err := strconv.ParseUint(args.ID, 10, 64)
	if err != nil {
		return nil
	}
	result := organizationservice.Find(id)
	o := organization{result}
	return &OrganizationResolver{&o}
}
