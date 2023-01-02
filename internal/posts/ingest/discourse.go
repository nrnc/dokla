package ingest

import (
	"encoding/json"
	"net/http"
	"time"
)

type Discourse struct {
	Name               string `json:"name,omitempty"`
	Username           string `json:"username,omitempty"`
	Blurb              string `json:"blurb,omitempty"`
	TopicTitleHeadline string `json:"topic_title_headline,omitempty"`
	CreatedAt          string `json:"created_at,omitempty"` //02-01-2006 same as twitter
	Avatar             string `json:"avatar,omitempty"`
	PostNumber         string `json:"post_number,omitempty"`
	PathParams         `json:"params,omitempty"`
	Meta               map[string]interface{} `json:"meta,omitempty"`
}

func (d *Discourse) SetPathParams(pathParams *PathParams) {
	d.App = pathParams.App
	d.Source = pathParams.Source
	d.Tenant = pathParams.Tenant
}

func (d *Discourse) Decode(req *http.Request) error {
	err := json.NewDecoder(req.Body).Decode(d)
	return err
}

func (d *Discourse) IsValid() bool {
	return d.PostNumber != ""
}

func (d *Discourse) SetTime() {
	if d.CreatedAt == "" {
		d.CreatedAt = time.Now().Format(time.RFC3339)
	} else {
		t, err := time.Parse("02-01-2006", d.CreatedAt)
		if err != nil {
			d.CreatedAt = time.Now().Format(time.RFC3339)
		} else {
			d.CreatedAt = t.Format(time.RFC3339)
		}
	}
}

func (t *Discourse) Adapt() *Request {
	nrequest := &Request{}
	nrequest.App = t.App
	nrequest.Tenant = t.Tenant
	nrequest.Source = t.Source
	nrequest.Name = t.Name
	nrequest.Username = t.Username
	nrequest.Avatar = t.Avatar
	nrequest.Content = t.Blurb
	nrequest.Meta = t.Meta
	nrequest.PostId = t.PostNumber
	nrequest.CreatedAt = t.CreatedAt
	nrequest.Title = t.TopicTitleHeadline

	return nrequest
}
