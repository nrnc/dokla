package ingest

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// newFetchPostsEndpoint creates an new endpoint that calls fetch service fetchposts
func newIngestEndpoint(f Service) endpoint.Endpoint {
	return func(
		cx context.Context,
		req interface{},
	) (res interface{}, err error) {

		r := req.(*Request)

		return f.Ingest(cx, r)
	}
}
