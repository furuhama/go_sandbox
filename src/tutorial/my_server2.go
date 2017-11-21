package tutorial

import (
	"fmt"
	"log"
	"myimage"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

// MyServer2 returns more details compared to MyServer
func MyServer2() {
	http.HandleFunc("/", handler2)
	http.HandleFunc("/lissajous", handler3)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler2 returns request header, host, remote addr, form
func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "Remote Addr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// counter returns how many times this server is called except '/counter'
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func handler3(w http.ResponseWriter, r *http.Request) {
	myimage.Lissajous(w)
}
