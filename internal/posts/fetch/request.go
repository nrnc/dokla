package fetch

import (
	"errors"
	"net/http"

	c "github.com/nrnc/dokla/internal/posts/fetch/consts"
	basehttp "github.com/unbxd/go-base/kit/transport/http"
)

type Request struct {
	Id     string
	Tenant string
	App    string
	After  string
	Before string
}

func GetRequest(req *http.Request) (*Request, error) {
	vars := basehttp.Parameters(req)
	app, ok := vars[c.APP]
	if !ok || app == "" {
		return nil, errors.New("error extracting app name")
	}
	tenant, ok := vars[c.TENANT]

	if !ok || tenant == "" {
		return nil, errors.New("error extracting tenant")
	}

	qparams := req.URL.Query()

	return &Request{
		Id:     qparams.Get("post_id"),
		App:    app,
		Tenant: tenant,
		After:  qparams.Get("after"),
		Before: qparams.Get("before"),
	}, nil
}
