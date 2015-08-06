// Package main is an entry point of the application.
package main

import (
	"flag"
	"net/http"
	"runtime"

	"github.com/anonx/sunplate/internal/skeleton/routes"

	"github.com/anonx/sunplate/log"
)

func main() {
	// Set max procs for multi-thread executing.
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Build the routes and handler.
	handler, err := routes.List.Build()
	if err != nil {
		log.Error.Fatal(err)
	}

	// Prepare a new server.
	s := &http.Server{
		Addr:    ":8080", // This is the default value of HTTP address.
		Handler: handler,
	}

	// Try to get some parameters from the received list of arguments,
	// e.g. "--addr=localhost:80".
	flag.StringVar(&s.Addr, "addr", s.Addr, "HTTP address that should be used by the app")
	flag.Parse()

	// Starting the server.
	log.Info.Printf(`Listening on "%s".`, s.Addr)
	log.Error.Fatal(serve(s))
}