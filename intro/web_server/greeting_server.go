package main

import (
	"fmt"
	"net/http"
)

// simple server runs on 8080 port, when go to http://localhost:8080/damla, you can see "Hello, damla!" text.
func main() {
	// each request calls greetingServer
	http.HandleFunc("/", greetingServer)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}

// greetingServer prints the Path component of the request URL r.
func greetingServer(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	if err != nil {
		return
	}
}
