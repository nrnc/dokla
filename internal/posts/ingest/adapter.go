package ingest

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	c "github.com/nrnc/dokla/internal/posts/ingest/consts"
)

type (
	PlayStore struct {
		Name       string `json:"name,omitempty"`
		GmailID    string `json:"mail,omitempty"`
		Review     string `json:"review,omitempty"`
		Title      string `json:"title,omitempty"`
		CreatedAt  string `json:"created_at,omitempty"` //2006-01-02
		Avatar     string `json:"avatar,omitempty"`
		ReviewID   string `json:"review_id,omitempty"`
		PathParams `json:"params,omitempty"`
		Meta       map[string]interface{} `json:"meta,omitempty"`
	}

	Twitter struct {
		Name       string `json:"name,omitempty"`
		Username   string `json:"username,omitempty"`
		Tweet      string `json:"tweet,omitempty"`
		Title      string `json:"title,omitempty"`
		CreatedAt  string `json:"created_at,omitempty"` //02-01-2006
		Avatar     string `json:"avatar,omitempty"`
		TweetId    string `json:"tweet_id,omitempty"`
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

func (p *PlayStore) SetPathParams(pathParams *PathParams) {
	p.App = pathParams.App
	p.Source = pathParams.Source
	p.Tenant = pathParams.Tenant
}

func (t *Twitter) SetPathParams(pathParams *PathParams) {
	t.App = pathParams.App
	t.Source = pathParams.Source
	t.Tenant = pathParams.Tenant
}

func PlayStoreAdaptor(req *http.Request, pathParams *PathParams) (*Request, error) {

	orequest := &PlayStore{}

	orequest.SetPathParams(pathParams)

	err := json.NewDecoder(req.Body).Decode(orequest)

	if err != nil {
		return nil, err
	}

	if orequest.ReviewID == "" {
		return nil, errors.New("review must contain review id")
	}

	if orequest.CreatedAt == "" {
		orequest.CreatedAt = time.Now().Format(time.RFC3339)
	} else {
		t, err := time.Parse("2006-01-02", orequest.CreatedAt)
		if err != nil {
			orequest.CreatedAt = time.Now().Format(time.RFC3339)
		} else {
			orequest.CreatedAt = t.Format(time.RFC3339)
		}
	}

	nrequest := &Request{}
	nrequest.App = orequest.App
	nrequest.Tenant = orequest.Tenant
	nrequest.Source = orequest.Source
	nrequest.Name = orequest.Name
	nrequest.Username = orequest.GmailID
	nrequest.Avatar = orequest.Avatar
	nrequest.Content = orequest.Review
	nrequest.Meta = orequest.Meta
	nrequest.PostId = orequest.ReviewID
	nrequest.CreatedAt = orequest.CreatedAt
	nrequest.Title = orequest.Title

	return nrequest, nil
}

func TwitterAdaptor(req *http.Request, pathParams *PathParams) (*Request, error) {

	orequest := &Twitter{}

	orequest.SetPathParams(pathParams)

	err := json.NewDecoder(req.Body).Decode(orequest)

	if err != nil {
		return nil, err
	}

	if orequest.TweetId == "" {
		return nil, errors.New("tweet must contain tweet id")
	}

	if orequest.CreatedAt == "" {
		orequest.CreatedAt = time.Now().Format(time.RFC3339)
	} else {
		t, err := time.Parse("02-01-2006", orequest.CreatedAt)
		if err != nil {
			orequest.CreatedAt = time.Now().Format(time.RFC3339)
		} else {
			orequest.CreatedAt = t.Format(time.RFC3339)
		}
	}

	nrequest := &Request{}
	nrequest.App = orequest.App
	nrequest.Tenant = orequest.Tenant
	nrequest.Source = orequest.Source
	nrequest.Name = orequest.Name
	nrequest.Username = orequest.Username
	nrequest.Avatar = orequest.Avatar
	nrequest.Content = orequest.Tweet
	nrequest.Meta = orequest.Meta
	nrequest.PostId = orequest.TweetId
	nrequest.CreatedAt = orequest.CreatedAt
	nrequest.Title = orequest.Title

	return nrequest, nil
}

func DiscourseAdaptor(req *http.Request, pathParams *PathParams) (*Request, error) {
	return nil, nil
}

func DefaultAdaptor(req *http.Request, pathParams *PathParams) (*Request, error) {
	return GetRequest(req, pathParams)
}
