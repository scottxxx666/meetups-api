package resolver

import (
	"strconv"

	"github.com/scottxxx666/meetups-api/model"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/scottxxx666/meetups-api/service/origanizationservice"
)

type origanization struct {
	model.Origanization
}

type OriganizationResolver struct {
	o *origanization
}

func (r *OriganizationResolver) ID() graphql.ID {
	id := strconv.FormatUint(r.o.ID, 10)
	return graphql.ID(id)
}

func (r *OriganizationResolver) Name() string {
	return r.o.Name
}

func (r *OriganizationResolver) Meetups() []*MeetupResolver {
	result := origanizationservice.GetMeetups(r.o.ID)
	var mr []*MeetupResolver
	for _, m := range result {
		mr = append(mr, &MeetupResolver{&meetup{m}})
	}
	return mr
}

func (r *Resolver) Origanization(args struct{ ID string }) *OriganizationResolver {
	id, err := strconv.ParseUint(args.ID, 10, 64)
	if err != nil {
		return nil
	}
	result := origanizationservice.Find(id)
	o := origanization{result}
	return &OriganizationResolver{&o}
}
