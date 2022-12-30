package ingest

import (
	"context"
	"net/http"
)

// decoder decodes the http request
func decoder(
	_ context.Context,
	req *http.Request,
) (interface{}, error) {

	pp, err := ExtractPathParams(req)

	if err != nil {
		return nil, err
	}

	request, err := AdaptorMap[pp.Source](req, pp)

	if err != nil {
		return nil, err
	}
	return request, nil
}
