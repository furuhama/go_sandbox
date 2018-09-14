// Package ch1 is for the chapter 1 of "The Go Programming Language"
package ch1

import (
	"fmt"
	"log"
	"net/http"
)

// MyServer receive request
// and send path element from request url as a response
func MyServer() {
	http.HandleFunc("/", handler) // handler is called by each request
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler returns path element from request url
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
