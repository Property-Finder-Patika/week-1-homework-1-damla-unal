package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

//the server runs the handler for each incoming request in a separate goroutine so that it can serve multiple requests simultaneously.
//However, if two concurrent requests try to update count at the same time, it might not be incremented consistently;
//the program would have a serious bug called a race condition.

// To avoid this problem, we must ensure that at most one goroutine accesses the variable at a time,
// which is the purpose of the mu.Lock() and mu.Unlock() calls that bracket each access of count.

func main() {
	http.HandleFunc("/path", handler)  // a request for /path calls handler
	http.HandleFunc("/count", counter) // a request for /count calls counter
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

// handler prints the path of the requested URL
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	_, err := fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	if err != nil {
		return
	}

}

// counter prints the number of request so far
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	_, err := fmt.Fprintf(w, "Request count: %d\n", count)
	if err != nil {
		return
	}
	mu.Unlock()
}
