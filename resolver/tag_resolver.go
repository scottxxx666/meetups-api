package resolver

import "github.com/scottxxx666/meetups-api/service"

// Tags resolve query tags
func (r *Resolver) Tags() []string {
	var result []string
	var ts service.TagService
	tags := ts.All()
	for _, t := range tags {
		result = append(result, t.ID)
	}
	return result
}
