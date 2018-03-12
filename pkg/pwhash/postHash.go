package pwhash

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

// TODO determine way to mock method calls (parseBodyForPw, hashAndB64Encode and timeSleep) for better encapsulated testing
func PostHashHandler(store Store, stats Stats) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		start := time.Now()
		if req.Method != "POST" {
			rw.WriteHeader(http.StatusNotFound)
			log.Printf("/hash http method %s is not supported", req.Method)
			return
		}

		pwInput, err := parseBodyForPw(req.Body)
		if err != nil {
			log.Printf("Error reading body: %v", err)
			http.Error(rw, "can't read body", http.StatusBadRequest)
			return
		}

		encoded := hashAndB64Encode(pwInput)
		//log.Printf("password form value='%s' hash='%s'", pwInput, encoded)
		key, err := store.Set(encoded)
		if err != nil {
			log.Printf("Error setting encoded value into the store: %v", err)
			http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		rw.WriteHeader(http.StatusOK)
		rw.Header().Set("Content-Type", "text/plain")

		io.WriteString(rw, key)
		duration := time.Now().Sub(start)
		stats.Time(duration.Nanoseconds() / int64(time.Microsecond))
		log.Printf("PostHashHandler complete in %s ", duration)
	})
}

/**
 * parse byte[] that looks like password=thePassword
 * return thePassword as a string
 */
func parseBodyForPw(bodyReader io.Reader) (string, error) {
	// TODO likely improvement is to avoid reading the entire buffer and conversion to strings
	body, err := ioutil.ReadAll(bodyReader)
	if err != nil {
		return "", err
	}

	parts := strings.SplitN(string(body), "=", 2)
	if len(parts) != 2 {
		return "", fmt.Errorf("body is invalid, missing '=' in '%s'", body)
	}
	return parts[1], nil
}

func hashAndB64Encode(input string) string {
	// TODO initializing hasher prior to endpoint call is likely an improvement
	hasher := sha512.New()
	hasher.Write([]byte(input))
	return base64.StdEncoding.EncodeToString(hasher.Sum(nil))
}
