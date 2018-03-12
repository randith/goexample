package main

import (
	"net/http"
	"github.com/randith/goexample/pkg/pwhash"
	"time"
	"log"
)

func main() {
	srv := startHttpServer()
	srv.ListenAndServe()

	// giving time after stop listening to finish up active requests
	time.Sleep(10 * time.Second)
	log.Printf("main: done. exiting")
}

func startHttpServer() *http.Server {
	log.Print("Starting server on port 8080")
	srv := &http.Server{Addr: ":8080"}

	http.Handle("/hash", pwhash.PostHashHandler())
	http.Handle("/shutdown", pwhash.PostShutdownHandler(srv))

	return srv
}