package resolver

import (
	"strconv"

	"github.com/scottxxx666/meetups-api/model"
	"github.com/scottxxx666/meetups-api/service/reviewservice"

	graphql "github.com/graph-gophers/graphql-go"
)

type review struct {
	model.Review
}

type ReviewResolver struct {
	r *review
}

func (r *ReviewResolver) ID() graphql.ID {
	id := strconv.FormatUint(r.r.ID, 10)
	return graphql.ID(id)
}

func (r *ReviewResolver) CreatedAt() string {
	return r.r.CreatedAt.String()
}

func (r *ReviewResolver) UpdatedAt() string {
	return r.r.UpdatedAt.String()
}

func (r *ReviewResolver) Desc() string {
	return r.r.Desc
}

func (r *ReviewResolver) Rating() int32 {
	return r.r.Rating
}

func (r *ReviewResolver) Member() *MemberResolver {
	return &MemberResolver{&member{r.r.Member}}
}

func (r *Resolver) Review(args struct{ ID string }) *ReviewResolver {
	id, err := strconv.ParseUint(args.ID, 10, 64)
	if err != nil {
		return nil
	}
	result := reviewservice.Find(id)
	review := review{result}
	return &ReviewResolver{&review}
}
