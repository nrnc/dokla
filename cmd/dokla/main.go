package main

import (
	"log"
	"os"
)

func main() {
	tp, err := NewDokla(
		withLogger(),
		withHttpTransport(),
		withMonitorHandler(),
		withFetchHandler(),
		withIngestHandler(),
	)

	if err != nil {
		log.Fatal("Error Starting dokla", err)
	}

	err = tp.Run(os.Args)
	if err != nil {
		tp.Close()
		log.Fatal("Error Starting dokla", err)
	}
}
