// Package tutorial is for the chapter 1 of "The Go Programming Language"
package tutorial

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// MyURLFetch parse url as args in command line
// then, send http request to url
func MyURLFetch() {
	for _, url := range os.Args[1:] {
		// if arg does not have "http://", add this phrase
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}
