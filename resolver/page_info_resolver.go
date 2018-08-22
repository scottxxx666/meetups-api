package resolver

import (
	"github.com/graph-gophers/graphql-go"
)

type PageInfoResolver struct {
	endCursor   *graphql.ID
	hasNextPage bool
}

func (r *PageInfoResolver) EndCursor() *graphql.ID {
	return r.endCursor
}

func (r *PageInfoResolver) HasNextPage() bool {
	return r.hasNextPage
}
