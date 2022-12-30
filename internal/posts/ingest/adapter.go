package ingest

import (
	"net/http"

	c "github.com/nrnc/dokla/internal/posts/ingest/consts"
)

type (
	PlayStore struct {
		Name       string `json:"name,omitempty"`
		GmailId    string `json:"mail,omitempty"`
		Review     string `json:"review,omitempty"`
		Title      string `json:"title,omitempty"`
		CreatedAt  string `json:"created_at,omitempty"`
		Avatar     string `json:"avatar,omitempty"`
		PostId     string `json:"post_id,omitempty"`
		PathParams `json:"params,omitempty"`
		Meta       map[string]interface{} `json:"meta,omitempty"`
	}
)

type Adaptor func(*http.Request, *PathParams) (*Request, error)

var AdaptorMap = map[string]Adaptor{
	c.PLAYSTORE: PlayStoreAdaptor,
	c.TWITTER:   TwitterAdaptor,
	c.DISCOURSE: DiscourseAdaptor,
	c.DEFAULT:   DefaultAdaptor,
}

func PlayStoreAdaptor(req *http.Request, pathParams *PathParams) (*Request, error) {
	return nil, nil
}

func TwitterAdaptor(req *http.Request, pathParams *PathParams) (*Request, error) {
	return nil, nil
}

func DiscourseAdaptor(req *http.Request, pathParams *PathParams) (*Request, error) {
	return nil, nil
}

func DefaultAdaptor(req *http.Request, pathParams *PathParams) (*Request, error) {
	return GetRequest(req, pathParams)
}
