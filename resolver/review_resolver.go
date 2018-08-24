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

// ReviewResolver resolve Review
type ReviewResolver struct {
	r *review
}

// ID resolver
func (r *ReviewResolver) ID() graphql.ID {
	id := strconv.FormatUint(r.r.ID, 10)
	return graphql.ID(id)
}

// CreatedAt resolver
func (r *ReviewResolver) CreatedAt() string {
	return r.r.CreatedAt.String()
}

// UpdatedAt resolver
func (r *ReviewResolver) UpdatedAt() string {
	return r.r.UpdatedAt.String()
}

// Desc resolver
func (r *ReviewResolver) Desc() string {
	return r.r.Desc
}

// Rating resolver
func (r *ReviewResolver) Rating() int32 {
	return r.r.Rating
}

// Member resolver
func (r *ReviewResolver) Member() *MemberResolver {
	return &MemberResolver{&member{r.r.Member}}
}

// Review resolve query Review
func (r *Resolver) Review(args struct{ ID string }) *ReviewResolver {
	id, err := strconv.ParseUint(args.ID, 10, 64)
	if err != nil {
		return nil
	}
	result := reviewservice.Find(id)
	review := review{result}
	return &ReviewResolver{&review}
}

// ReviewsConnectionResolver resolve ReviewConnection
type ReviewsConnectionResolver struct {
	reviews    []*review
	totalCount int32
}

// ReviewsConnection resolve query ReviewsConnection
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
		cursor, err := unmarshalReviewCursor(*args.After)
		if err != nil {
			panic(err)
		}
		result, count = reviewservice.GetByMeetupAfter(mid, int(args.First), cursor.getID(), cursor.getUpdatedAt())
	}
	for _, m := range result {
		rs = append(rs, &review{m})
	}
	return &ReviewsConnectionResolver{rs, count}
}

// TotalCount resolver
func (r *ReviewsConnectionResolver) TotalCount() int32 {
	return r.totalCount
}

// Edges resolver
func (r *ReviewsConnectionResolver) Edges() []*ReviewEdgeResolver {
	var resolvers []*ReviewEdgeResolver
	for _, re := range r.reviews {
		resolvers = append(resolvers, &ReviewEdgeResolver{re})
	}
	return resolvers
}

// PageInfo resolver
func (r *ReviewsConnectionResolver) PageInfo() *PageInfoResolver {
	e := r.Edges()
	if len(e) == 0 {
		return &PageInfoResolver{nil, false}
	}
	re := e[len(e)-1]
	c := re.Cursor()
	return &PageInfoResolver{&c, int(r.totalCount) != len(e)}
}

// ReviewEdgeResolver resolve ReviewEdge
type ReviewEdgeResolver struct {
	review *review
}

// Node resolver
func (r *ReviewEdgeResolver) Node() *ReviewResolver {
	return &ReviewResolver{r.review}
}

// Cursor resolver
func (r *ReviewEdgeResolver) Cursor() graphql.ID {
	return marshalReviewCursor(r.review.ID, r.review.UpdatedAt)
}

type reviewCursor struct {
	ID       string
	UpdateAt string
}

func (r *reviewCursor) getID() uint64 {
	id, err := strconv.ParseUint(r.ID, 10, 64)
	if err != nil {
		panic(nil)
	}
	return id
}

func (r *reviewCursor) getUpdatedAt() string {
	return r.UpdateAt
}

func marshalReviewCursor(id uint64, updateAt time.Time) graphql.ID {
	m := reviewCursor{strconv.FormatUint(id, 10), updateAt.String()}
	j, err := json.Marshal(m)
	if err != nil {
		panic("Json Marshal error")
	}
	return graphql.ID(base64.URLEncoding.EncodeToString([]byte(j)))
}

func unmarshalReviewCursor(a string) (*reviewCursor, error) {
	s, err := base64.URLEncoding.DecodeString(a)
	if err != nil {
		return nil, err
	}
	cursor := reviewCursor{}
	byt := []byte(s)
	if err := json.Unmarshal(byt, &cursor); err != nil {
		panic(err)
	}
	return &cursor, nil
}
