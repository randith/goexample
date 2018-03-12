package pwhash

import (
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

func GetHashHandler(store Store) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		start := time.Now()
		if req.Method != "GET" {
			rw.WriteHeader(http.StatusNotFound)
			log.Printf("/hash/{id} http method %s is not supported", req.Method)
			return
		}

		path := req.URL.Path
		parts := strings.SplitN(string(path), "/", 3)
		if len(parts) != 3 {
			http.Error(rw, "url path is invalid, missing 'id' for password hash", http.StatusBadRequest)
			return
		}

		hash, err := store.Get(parts[2])
		if err != nil {
			http.Error(rw, "Not Found", http.StatusNotFound)
			return
		}

		rw.WriteHeader(http.StatusOK)
		rw.Header().Set("Content-Type", "text/plain")
		io.WriteString(rw, hash)

		log.Printf("GetHashHandler complete in %s", time.Now().Sub(start))
	})
}
