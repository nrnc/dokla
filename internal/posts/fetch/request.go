package fetch

import (
	"errors"
	"net/http"

	c "github.com/nrnc/dokla/internal/posts/fetch/consts"
	basehttp "github.com/unbxd/go-base/kit/transport/http"
)

type Request struct {
	Id     string `json:"id"`
	App    string `json:"app"`
	After  string `json:"after"`
	Before string `json:"before"`
}

func GetRequest(req *http.Request) (*Request, error) {
	vars := basehttp.Parameters(req)
	app, ok := vars[c.APP]
	if !ok || app == "" {
		return nil, errors.New("error extracting app name")
	}
	qparams := req.URL.Query()

	return &Request{
		Id:     qparams.Get("id"),
		App:    app,
		After:  qparams.Get("after"),
		Before: qparams.Get("before"),
	}, nil
}
