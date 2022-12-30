package ingest

import (
	nethttp "net/http"

	"github.com/unbxd/go-base/kit/transport/http"

	"github.com/unbxd/go-base/utils/log"
)

type (
	// IngestHandler is the main handler struct
	IngestHandler struct {
		logger log.Logger
	}

	HandlerOption func(*IngestHandler)
)

// HandlerWithLogger adds the logger
func HandlerWithLogger(logger log.Logger) HandlerOption {
	return func(sh *IngestHandler) {
		sh.logger = logger
	}
}

// NewIngestHandler returns a new fetch handler object initialized from the options
func NewIngestHandler(opts ...HandlerOption) *IngestHandler {
	sh := &IngestHandler{}
	for _, opt := range opts {
		opt(sh)
	}
	return sh
}

// Bind registers the handlers
func Bind(transport *http.Transport, http_handler nethttp.Handler) {
	transport.Mux().Handler(
		nethttp.MethodPut,
		"/:source/:tenant/:app/posts",
		http_handler,
	)
}

// ingestHttpHandler creates http.Handler init with fetchposts endpoint
func (ih *IngestHandler) ingestHttpHandler(repo Repository) http.Handler {
	svc, _ := NewService(
		ih.logger,
		repo,
	)
	ep := newIngestEndpoint(svc)
	return http.Handler(ep)
}

// HTTPHandler returns an http handler with options and fetch endpoint
func (ih *IngestHandler) HTTPHandler(
	repo Repository,
	opts ...http.HandlerOption) nethttp.Handler {
	options := append(
		[]http.HandlerOption{
			http.HandlerWithEncoder(encoder),
			http.HandlerWithErrorEncoder(errorEncoder),
			http.HandlerWithDecoder(decoder),
			http.HandlerWithFilter(http.PanicRecovery(ih.logger)),
		}, opts...,
	)
	return http.NewHandler(
		ih.ingestHttpHandler(repo),
		options...)
}
