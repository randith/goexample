package main

import (
	"net/http"
	"log"
	"github.com/randith/goexample/pkg/pwhash"
)

func main() {
	mux := http.NewServeMux()
	// get the value of a key
	mux.HandleFunc("/hash", pwhash.PostHashHandler)
	// set the value of a key

	log.Printf("starting server on port 8080")

	// http.ListenAndServe takes in an http.Handler as its second parameter.
	// since ServeMux implements a ServeHTTP function, it is also an http.Handler,
	// so we can pass it here.
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}