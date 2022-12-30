package ingest

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	c "github.com/nrnc/dokla/internal/posts/ingest/consts"
	basehttp "github.com/unbxd/go-base/kit/transport/http"
)

type (
	Request struct {
		Name       string `json:"name,omitempty" bson:"name,omitempty"`
		Username   string `json:"username,omitempty" bson:"username,omitempty"`
		Content    string `json:"content,omitempty" bson:"content,omitempty"`
		Title      string `json:"title,omitempty" bson:"title,omitempty"`
		CreatedAt  string `json:"created_at,omitempty" bson:"created_at,omitempty"`
		Avatar     string `json:"avatar,omitempty" bson:"avatar,omitempty"`
		PostId     string `json:"post_id,omitempty" bson:"post_id,omitempty"`
		PathParams `json:"params,omitempty" bson:"inline"`
	}
	PathParams struct {
		App    string `json:"app,omitempty" bson:"app,omitempty"`
		Tenant string `json:"tenant,omitempty" bson:"tenant,omitempty"`
		Source string `json:"source,omitempty" bson:"source,omitempty"`
	}
)

func NewRequest() *Request {
	return &Request{
		CreatedAt: time.Now().Format(time.RFC3339),
	}
}

func ExtractPathParams(r *http.Request) (*PathParams, error) {
	// Read Mux's Vars
	vars := basehttp.Parameters(r)
	app, ok := vars[c.APP]

	if !ok {
		return nil, errors.New("Error extracting app from the request")
	}

	tenant, ok := vars[c.TENANT]
	if !ok {
		return nil, errors.New("Error extracting tenant from the request")
	}

	source, ok := vars[c.SOURCE]
	if !ok {
		return nil, errors.New("Error extracting source from the request")
	}

	return &PathParams{
		app, tenant, source,
	}, nil
}

func GetRequest(req *http.Request) (*Request, error) {
	pp, err := ExtractPathParams(req)

	if err != nil {
		return nil, err
	}

	request := NewRequest()
	request.App = pp.App
	request.Tenant = pp.Tenant
	request.Source = pp.Source
	err = json.NewDecoder(req.Body).Decode(request)
	if err != nil {
		return nil, err
	}

	return request, nil
}
