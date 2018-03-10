package pwhash

import (
	"testing"
	"fmt"
	"net/http/httptest"
	"net/http"
	"io"
)

type hashAndEncodeTest struct {
	in string
	out string
}

var hashAndEncodeTests = []hashAndEncodeTest {
	{ "angryMonkey", "ZEHhWB65gUlzdVwtDQArEyx+KVLzp/aTaRaPlBzYRIFj6vjFdqEb0Q5B8zVKCZ0vKbZPZklJz0Fd7su2A+gf7Q==" },
	{ "something else", "NFJPgMVW1OXAHj7S0GASZTwC1DbKo++ACCVQFfne/x8A6KcK42g9BtGbXcn7TBMPqkVZ+wrRygXGPpLmFuJJ+A==" },
	{ "", "z4PhNX7vuL3xVChQ1m2AB9Yg5AULVxXcg/SpIdNs6c5H0NE8XYXysP+DGNKHfuwvY7kxvUdBeoGlODJ6+SfaPg==" },
}

type Reader struct {
	read string
	done bool
}

func NewReader(toRead string) *Reader {
	return &Reader{toRead, false}
}

func (r *Reader) Read(p []byte) (n int, err error) {
	if r.done {
		return 0, io.EOF
	}
	for i, b := range []byte(r.read) {
		p[i] = b
	}
	r.done = true
	return len(r.read), nil
}

func Test_hashAndEncodeTests(t *testing.T)  {
	for _, test := range hashAndEncodeTests {
		actual := hashAndB64Encode(test.in)
		if actual != test.out {
			t.Fatal(fmt.Sprintf("test in '%s' has actual of '%s' instead of expected '%s'", test.in, actual, test.out))
		}
	}
}

// TODO 5 seconds for this test is way too long, need to refactor to either mock something that does the sleep or send in a very short sleep
func Test_PostHashHandler(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("POST", "/hash", NewReader("password=" + hashAndEncodeTests[0].in))
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(PostHashHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := hashAndEncodeTests[0].out
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: actual='%v' expected='%v'",
			rr.Body.String(), expected)
	}
}