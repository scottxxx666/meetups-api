package resolver

import (
	"strconv"

	"github.com/scottxxx666/meetups-api/service"

	"github.com/graph-gophers/graphql-go"

	"github.com/scottxxx666/meetups-api/model"
)

type location struct {
	model.Location
}

// LocationResolver resolve Location
type LocationResolver struct {
	l *location
}

// ID resolve ID
func (r *LocationResolver) ID() graphql.ID {
	id := strconv.FormatUint(r.l.ID, 10)
	return graphql.ID(id)
}

// Name resolve Name
func (r *LocationResolver) Name() string {
	return r.l.Name
}

// Locations resolve query Locations
func (r *Resolver) Locations() []*LocationResolver {
	var lrs []*LocationResolver
	var s service.LocationService
	locations := s.All()
	for _, l := range locations {
		lrs = append(lrs, &LocationResolver{&location{l}})
	}
	return lrs
}
