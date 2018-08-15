package resolver

import (
	"strconv"

	"github.com/scottxxx666/meetups-api/model"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/scottxxx666/meetups-api/service/memberservice"
)

type member struct {
	model.Member
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

func (r *Resolver) Member(args struct{ ID string }) *MemberResolver {
	id, err := strconv.ParseUint(args.ID, 10, 64)
	if err != nil {
		return nil
	}
	result := memberservice.Find(id)
	m := member{result}
	return &MemberResolver{&m}
}
