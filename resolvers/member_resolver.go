package resolvers

import (
	"strconv"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/scottxxx666/meetups-api/services/memberservice"
)

type member struct {
	ID   uint64
	Name string
}

type MemberResolver struct {
	m *member
}

func (r *MemberResolver) ID() graphql.ID {
	id := strconv.FormatUint(r.m.ID, 10)
	return graphql.ID(id)
}

func (r *MemberResolver) Name() string {
	return r.m.Name
}

// func (r *Resolver) Member(args struct{ ID graphql.ID }) *MemberResolver {
func (r *Resolver) Member(args struct{ ID string }) *MemberResolver {
	id, err := strconv.ParseUint(args.ID, 10, 64)
	if err != nil {
		return nil
	}
	model := memberservice.Find(id)
	m := member{model.ID, model.Name}
	return &MemberResolver{&m}
}
