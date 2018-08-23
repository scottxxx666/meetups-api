package resolver

import (
	"encoding/base64"
	"encoding/json"
	"strconv"
	"time"

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

type ReviewsConnectionResolver struct {
	reviews    []*review
	totalCount int32
}

func (r *Resolver) ReviewsConnection(args connectionArgs) *ReviewsConnectionResolver {
	mid, err := strconv.ParseUint(args.MeetupID, 10, 64)
	if err != nil {
		return nil
	}
	var rs []*review
	var result []model.Review
	var count int32
	if args.After == nil {
		result, count = reviewservice.GetByMeetup(mid, int(args.First))
	} else {
		s, err := base64.URLEncoding.DecodeString(*args.After)
		if err != nil {
			panic(err)
		}
		var after map[string]string
		byt := []byte(s)
		if err := json.Unmarshal(byt, &after); err != nil {
			panic(err)
		}
		id, err := strconv.ParseUint(after["id"], 10, 64)
		if err != nil {
			return nil
		}
		updatedAt := after["updated_at"]
		result, count = reviewservice.GetByMeetupAfter(mid, int(args.First), id, updatedAt)
	}
	for _, m := range result {
		rs = append(rs, &review{m})
	}
	return &ReviewsConnectionResolver{rs, count}
}

func (r *ReviewsConnectionResolver) TotalCount() int32 {
	return r.totalCount
}

func (r *ReviewsConnectionResolver) Edges() []*ReviewEdgeResolver {
	var resolvers []*ReviewEdgeResolver
	for _, re := range r.reviews {
		resolvers = append(resolvers, &ReviewEdgeResolver{re})
	}
	return resolvers
}

func (r *ReviewsConnectionResolver) PageInfo() *PageInfoResolver {
	e := r.Edges()
	if len(e) == 0 {
		return &PageInfoResolver{nil, false}
	}
	re := e[len(e)-1]
	c := re.Cursor()
	return &PageInfoResolver{&c, int(r.totalCount) != len(e)}
}

type ReviewEdgeResolver struct {
	review *review
}

func (r *ReviewEdgeResolver) Node() *ReviewResolver {
	return &ReviewResolver{r.review}
}

func (r *ReviewEdgeResolver) Cursor() graphql.ID {
	return marshalCursor(r.review.ID, r.review.UpdatedAt)
}

func marshalCursor(id uint64, updateAt time.Time) graphql.ID {
	m := map[string]string{"id": strconv.FormatUint(id, 10), "updated_at": updateAt.String()}
	j, err := json.Marshal(m)
	if err != nil {
		panic("Json Marshal error")
	}
	return graphql.ID(base64.URLEncoding.EncodeToString([]byte(j)))
}
