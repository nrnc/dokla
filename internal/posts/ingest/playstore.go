package ingest

import (
	"encoding/json"
	"net/http"
	"time"
)

type PlayStore struct {
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

func (p *PlayStore) SetPathParams(pathParams *PathParams) {
	p.App = pathParams.App
	p.Source = pathParams.Source
	p.Tenant = pathParams.Tenant
}

func (p *PlayStore) Decode(req *http.Request) error {
	err := json.NewDecoder(req.Body).Decode(p)
	return err
}

func (p *PlayStore) IsValid() bool {
	return p.ReviewID != ""
}

func (p *PlayStore) SetTime() {
	if p.CreatedAt == "" {
		p.CreatedAt = time.Now().Format(time.RFC3339)
	} else {
		t, err := time.Parse("2006-01-02", p.CreatedAt)
		if err != nil {
			p.CreatedAt = time.Now().Format(time.RFC3339)
		} else {
			p.CreatedAt = t.Format(time.RFC3339)
		}
	}
}

func (p *PlayStore) Adapt() *Request {

	nrequest := &Request{}
	nrequest.App = p.App
	nrequest.Tenant = p.Tenant
	nrequest.Source = p.Source
	nrequest.Name = p.Name
	nrequest.Username = p.GmailID
	nrequest.Avatar = p.Avatar
	nrequest.Content = p.Review
	nrequest.Meta = p.Meta
	nrequest.PostId = p.ReviewID
	nrequest.CreatedAt = p.CreatedAt
	nrequest.Title = p.Title

	return nrequest
}
