package pwhash

import (
	"time"
	"net/http"
	"log"
	"encoding/json"
)

func GetStatsHandler(stats Stats) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		start := time.Now()
		if req.Method != "GET" {
			rw.WriteHeader(http.StatusNotFound)
			log.Printf("/stats http method %s is not supported", req.Method)
			return
		}

		stats.Get()

		rw.WriteHeader(http.StatusOK)
		json.NewEncoder(rw).Encode(stats.Get())
		// TODO bug: seems to still be returning text/plain content type
		rw.Header().Set("Content-Type", "application/json")
		log.Printf("GetStatsHandler complete in %s ", time.Now().Sub(start))
	})
}