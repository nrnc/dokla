package main

import (
	"github.com/nrnc/dokla/cmd/dokla/flags"
	"github.com/nrnc/dokla/internal/posts/fetch"
	"github.com/nrnc/dokla/internal/posts/ingest"
	"github.com/nrnc/dokla/internal/posts/monitor"
	"github.com/nrnc/dokla/internal/posts/utils"
	"github.com/pkg/errors"
	"github.com/unbxd/go-base/kit/transport/http"
	"github.com/unbxd/go-base/utils/log"
)

// init dokla logger
func withLogger() Option {
	return func(dokla *Dokla) (err error) {
		logger, err := log.NewZapLogger(
			log.ZapWithLevel(flags.LogLevel),
			log.ZapWithEncoding(flags.LogEncoding),
			log.ZapWithOutput([]string{flags.LogOutput}),
		)

		if err != nil {
			return errors.Wrap(err, "create logger failed")
		}
		dokla.logger = logger
		return
	}
}

func withMongoClient() Option {
	return func(dokla *Dokla) (err error) {
		client, err := utils.MongoClient(flags.DbConn, dokla.logger)

		if err != nil {
			return errors.Wrap(err, "create mongo connection failed")
		}
		dokla.mongoClient = client
		return
	}
}

// Init dokla http transport with a few defaults
func withHttpTransport() Option {
	return func(dokla *Dokla) (err error) {
		opts := []http.TransportOption{
			http.WithLogger(dokla.logger),
			http.TransportWithFilter(http.PanicRecovery(dokla.logger)),
			http.WithFullDefaults(),
		}

		tr, err := http.NewTransport(
			flags.HTTPHost,
			flags.HTTPPort,
			opts...,
		)
		if err != nil {
			return errors.Wrap(err, "create server failed")
		}

		dokla.httpTr = tr
		return
	}
}

// Add a monitor API. used for liveness and readiness probes
func withMonitorHandler() Option {
	return func(dokla *Dokla) (err error) {
		dokla.httpTr.Get("/monitor", monitor.MonitorHandlerFn())
		return
	}
}

func withIngestHandler() Option {
	return func(d *Dokla) error {
		ih := ingest.NewIngestHandler(
			ingest.HandlerWithLogger(d.logger),
		)
		mr := ingest.NewMongoIngestRepo(d.mongoClient)
		ingest.Bind(d.httpTr, ih.HTTPHandler(mr))

		return nil
	}
}

func withFetchHandler() Option {
	return func(d *Dokla) error {
		fh := fetch.NewFetchHandler(
			fetch.HandlerWithLogger(d.logger),
		)
		fetch.Bind(d.httpTr, fh.HTTPHandler(d.mongoClient))

		return nil
	}
}
