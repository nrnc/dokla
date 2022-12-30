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

	request, err := GetRequest(req)
	if err != nil {
		return nil, err
	}
	return request, nil
}
