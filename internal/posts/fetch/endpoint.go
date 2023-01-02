package fetch

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// newFetchPostsEndpoint creates an new endpoint that calls fetch service interface
func newFetchPostsEndpoint(f Service) endpoint.Endpoint {
	return func(
		cx context.Context,
		req interface{},
	) (res interface{}, err error) {

		r := req.(*Request)
		if r.Id != "" {
			return f.FetchById(cx, r)
		}

		return f.FetchByDuration(cx, r)
	}
}
