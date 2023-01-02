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

	r, err := BuildRequest(req, pp)

	if err != nil {
		return nil, err
	}
	return r, nil
}
