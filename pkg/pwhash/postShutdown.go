package pwhash

import (
	"context"
	"log"
	"net/http"
	"time"
)

func PostShutdownHandler(srv *http.Server) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		start := time.Now()
		if req.Method != "POST" {
			rw.WriteHeader(http.StatusNotFound)
			log.Printf("/shutdown http method %s is not supported", req.Method)
			return
		}

		log.Printf("PostShutdownHandler issuing server shutdown")
		ctx, _ := context.WithTimeout(context.Background(), 9*time.Second)

		srv.Shutdown(ctx)

		log.Printf("PostShutdownHandler complete in %s", time.Now().Sub(start))
	})
}
