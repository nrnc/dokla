package main

import (
	"fmt"
	cl "log"
	"os"
	"os/signal"

	"github.com/nrnc/dokla/cmd/dokla/flags"
	"github.com/unbxd/go-base/kit/transport/http"
	"github.com/unbxd/go-base/utils/log"
	"github.com/urfave/cli/v2"
)

type (
	// Dokla Options
	Option func(*Dokla) error

	Dokla struct {
		*cli.App
		logger log.Logger
		httpTr *http.Transport
	}
)

// Listens for os.Signal on an int channel and starts the Transport
func (t *Dokla) Open() (err error) {
	var (
		intCh = make(chan os.Signal, 1)
	)

	// start the http transport in a go routine
	go func() {
		t.logger.Info("Starting Dokla Server", log.String("addr", t.httpTr.Addr))
		er := t.httpTr.Open()
		if er != nil {
			cl.Fatal(er)
		}
	}()

	signal.Notify(
		intCh, os.Interrupt,
	)

	// monitor for OS Interrupt and close the application gracefully
	for range intCh {
		fmt.Println("Received OS interrupt. Shutting down DOKLA")
		t.Close()
		os.Exit(0)

	}

	return nil
}

// Shutdown the Transport
func (t *Dokla) Close() {
	if t.httpTr != nil {
		_ = t.httpTr.Close()
	}
}

// Create a new Dokla application
func NewDokla(options ...Option) (*Dokla, error) {
	dokla := &Dokla{
		App: &cli.App{
			Name:  ccName,
			Usage: ccUsage,
			Flags: flags.Flags(),
			Before: func(cx *cli.Context) (err error) {
				return
			},
		},
	}

	startcmd := &cli.Command{
		Name:    startName,
		Aliases: []string{startAlias},
		Usage:   startUsage,
		Before: func(cx *cli.Context) (err error) {
			for _, o := range options {
				e := o(dokla)
				if e != nil {
					return cli.Exit(
						"Error starting Dokla: [ "+e.Error()+" ]", 9,
					)
				}
			}
			return
		},

		Action: func(cx *cli.Context) (err error) {
			err = dokla.Open()
			if err != nil {
				return cli.Exit(
					err, 9,
				)
			}

			return
		},
	}

	dokla.Commands = append(
		dokla.Commands, startcmd,
	)

	return dokla, nil
}
