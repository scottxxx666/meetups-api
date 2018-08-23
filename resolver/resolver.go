package resolver

type Resolver struct{}

type connectionArgs struct {
	MeetupID string
	First    float64
	After    *string
}
