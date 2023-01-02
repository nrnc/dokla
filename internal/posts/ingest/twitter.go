package ingest

import (
	"encoding/json"
	"net/http"
	"time"
)

type Twitter struct {
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

func (t *Twitter) SetPathParams(pathParams *PathParams) {
	t.App = pathParams.App
	t.Source = pathParams.Source
	t.Tenant = pathParams.Tenant
}

func (t *Twitter) Decode(req *http.Request) error {
	err := json.NewDecoder(req.Body).Decode(t)
	return err
}

func (t *Twitter) IsValid() bool {
	return t.TweetId != ""
}

func (t *Twitter) SetTime() {
	if t.CreatedAt == "" {
		t.CreatedAt = time.Now().Format(time.RFC3339)
	} else {
		ti, err := time.Parse("02-01-2006", t.CreatedAt)
		if err != nil {
			t.CreatedAt = time.Now().Format(time.RFC3339)
		} else {
			t.CreatedAt = ti.Format(time.RFC3339)
		}
	}
}

func (t *Twitter) Adapt() *Request {
	nrequest := &Request{}
	nrequest.App = t.App
	nrequest.Tenant = t.Tenant
	nrequest.Source = t.Source
	nrequest.Name = t.Name
	nrequest.Username = t.Username
	nrequest.Avatar = t.Avatar
	nrequest.Content = t.Tweet
	nrequest.Meta = t.Meta
	nrequest.PostId = t.TweetId
	nrequest.CreatedAt = t.CreatedAt
	nrequest.Title = t.Title
	return nrequest
}
