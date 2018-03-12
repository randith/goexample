package main

import (
	"github.com/randith/goexample/pkg/pwhash"
	"log"
	"net/http"
	"time"
)

func main() {
	store := pwhash.NewInMemoryStore()
	statsStore := pwhash.NewInMemoryStats()
	srv := startHttpServer(store, statsStore)
	srv.ListenAndServe()

	// giving time after stop listening to finish up active requests
	time.Sleep(10 * time.Second)
	log.Printf("main: done. exiting")
}

func startHttpServer(store pwhash.Store, statsStore pwhash.Stats) *http.Server {
	log.Print("Starting server on port 8080")
	srv := &http.Server{Addr: ":8080"}

	http.Handle("/hash", pwhash.PostHashHandler(store, statsStore))
	http.Handle("/hash/", pwhash.GetHashHandler(store))
	http.Handle("/shutdown", pwhash.PostShutdownHandler(srv))
	http.Handle("/stats", pwhash.GetStatsHandler(statsStore))

	return srv
}
