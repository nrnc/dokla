package monitor

import (
	"context"
	nethttp "net/http"

	"github.com/unbxd/go-base/kit/transport/http"
)

var monitorMessage = "want a dokla? let's go to Gujarat!!"

func MonitorHandlerFn() http.HandlerFunc {
	return func(
		_ context.Context, req *nethttp.Request,
	) (*nethttp.Response, error) {
		return http.NewResponse(
			req, http.ResponseWithBytes([]byte(monitorMessage)),
		), nil
	}
}
