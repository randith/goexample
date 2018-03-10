package pwhash

import (
	"net/http"
	"log"
	"encoding/base64"
	"crypto/sha512"
	"io"
	"time"
	"io/ioutil"
	"strings"
	"fmt"
)

// TODO determine way to mock method calls (parseBodyForPw, hashAndB64Encode and timeSleep) for better encapsulated testing
func PostHashHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	if r.Method != "POST" {
		w.WriteHeader(http.StatusNotFound)
		log.Printf("/hash http method %s is not supported", r.Method)
		return
	}

	pwInput, err := parseBodyForPw(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	encoded := hashAndB64Encode(pwInput)
	//log.Printf("password form value='%s' hash='%s'", pwInput, encoded)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")

	io.WriteString(w, encoded)
	time.Sleep(5 * time.Second)
	log.Printf("PostHashHandler complete in %s ", time.Now().Sub(start))
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

