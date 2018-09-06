package resolver

import (
	"strconv"

	"github.com/scottxxx666/meetups-api/model"
	"github.com/scottxxx666/meetups-api/service"

	graphql "github.com/graph-gophers/graphql-go"
)

type member struct {
	model.Member
}

// MemberResolver resolve Member
type MemberResolver struct {
	m *member
}

// ID resolver
func (r *MemberResolver) ID() graphql.ID {
	id := strconv.FormatUint(r.m.ID, 10)
	return graphql.ID(id)
}

// Name resolver
func (r *MemberResolver) Name() string {
	return r.m.Name
}

// Member resolve query Member
func (r *Resolver) Member(args struct{ ID string }) *MemberResolver {
	id, err := strconv.ParseUint(args.ID, 10, 64)
	if err != nil {
		return nil
	}
	var ms service.MemberService
	result := ms.Find(id)
	m := member{result}
	return &MemberResolver{&m}
}
