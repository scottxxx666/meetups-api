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

type OrganizationResolver struct {
	o *organization
}

func (r *OrganizationResolver) ID() graphql.ID {
	id := strconv.FormatUint(r.o.ID, 10)
	return graphql.ID(id)
}

func (r *OrganizationResolver) Name() string {
	return r.o.Name
}

func (r *OrganizationResolver) Meetups() []*MeetupResolver {
	result := organizationservice.GetMeetups(r.o.ID)
	var mr []*MeetupResolver
	for _, m := range result {
		mr = append(mr, &MeetupResolver{&meetup{m}})
	}
	return mr
}

func (r *Resolver) Organization(args struct{ ID string }) *OrganizationResolver {
	id, err := strconv.ParseUint(args.ID, 10, 64)
	if err != nil {
		return nil
	}
	result := organizationservice.Find(id)
	o := organization{result}
	return &OrganizationResolver{&o}
}
