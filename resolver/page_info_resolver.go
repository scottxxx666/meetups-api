package resolver

import (
	"github.com/graph-gophers/graphql-go"
)

// PageInfoResolver resolve PageInfo
type PageInfoResolver struct {
	endCursor   *graphql.ID
	hasNextPage bool
}

// EndCursor resolver
func (r *PageInfoResolver) EndCursor() *graphql.ID {
	return r.endCursor
}

// HasNextPage resolver
func (r *PageInfoResolver) HasNextPage() bool {
	return r.hasNextPage
}
