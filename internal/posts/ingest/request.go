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
		Meta       map[string]interface{} `json:"meta" bson:"meta,omitempty"`
	}
	PathParams struct {
		App    string `json:"app,omitempty" bson:"app,omitempty"`
		Tenant string `json:"tenant,omitempty" bson:"tenant,omitempty"`
		Source string `json:"source,omitempty" bson:"source,omitempty"`
	}

	Meta struct {
		AppVersion string `json:"app_version" bson:"app_version,omitempty"`
		Device     string `json:"Device" bson:"device,omitempty"`
		Location   string `json:"location" bson:"location,omitempty"`
		Language   string `json:"language" bson:"language,omitempty"`
	}
)

func NewRequest() *Request {
	return &Request{
		CreatedAt: time.Now().Format(time.RFC3339),
	}
}

func (r *Request) SetPathParams(pathParams *PathParams) {
	r.App = pathParams.App
	r.Source = pathParams.Source
	r.Tenant = pathParams.Tenant
}

func (r *Request) Decode(req *http.Request) error {
	err := json.NewDecoder(req.Body).Decode(r)
	return err
}

func (r *Request) SetTime() {
	if r.CreatedAt == "" {
		r.CreatedAt = time.Now().Format(time.RFC3339)
	}
}

func (r *Request) Adapt() *Request {
	return r
}

func (r *Request) IsValid() bool {
	return r.PostId != ""
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
