package resolver

// Resolver resolve queries and mutations
type Resolver struct{}

type connectionArgs struct {
	MeetupID string
	First    float64
	After    *string
}
