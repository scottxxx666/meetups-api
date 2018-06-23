package resolvers

import graphql "github.com/graph-gophers/graphql-go"

type person struct {
	FirstName graphql.ID
}

type personResolver struct {
	p *person
}

func (r *personResolver) ID() graphql.ID {
	return r.p.FirstName
}

func (r *Resolver) Person(args struct{ ID graphql.ID }) *personResolver {
	p := person{"AAA"}
	return &personResolver{&p}
}
