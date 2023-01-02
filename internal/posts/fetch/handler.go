package fetch

import (
	nethttp "net/http"

	"github.com/unbxd/go-base/kit/transport/http"

	"github.com/unbxd/go-base/utils/log"
)

type (
	// FetchHandler is the main fetch handler struct
	FetchHandler struct {
		logger log.Logger
	}

	HandlerOption func(*FetchHandler)
)

// adds the logger
func HandlerWithLogger(logger log.Logger) HandlerOption {
	return func(sh *FetchHandler) {
		sh.logger = logger
	}
}

// NewFetchHandler returns a new fetch handler object initialized from the options
func NewFetchHandler(opts ...HandlerOption) *FetchHandler {
	sh := &FetchHandler{}
	for _, opt := range opts {
		opt(sh)
	}
	return sh
}

// Bind registers the handlers
func Bind(transport *http.Transport, http_handler nethttp.Handler) {
	transport.Mux().Handler(
		nethttp.MethodGet,
		"/:tenant/:app/posts",
		http_handler,
	)
}

// fetchHttpHandler creates http.Handler init with fetchposts endpoint
func (fh *FetchHandler) fetchHttpHandler(repo Repository) http.Handler {
	svc, _ := NewService(
		fh.logger,
		repo,
	)
	ep := newFetchPostsEndpoint(svc)
	return http.Handler(ep)
}

// HTTPHandler returns an http handler with options and fetch endpoint
func (fh *FetchHandler) HTTPHandler(
	repo Repository,
	opts ...http.HandlerOption) nethttp.Handler {
	options := append(
		[]http.HandlerOption{
			http.HandlerWithEncoder(encoder),
			http.HandlerWithErrorEncoder(errorEncoder),
			http.HandlerWithDecoder(decoder),
			http.HandlerWithFilter(http.PanicRecovery(fh.logger)),
		}, opts...,
	)
	return http.NewHandler(
		fh.fetchHttpHandler(repo),
		options...)
}
