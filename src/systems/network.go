// Package systems is for system layor program
package systems

import (
	"io"
	"net"
	"net/http"
	"os"
)

// TCPConnect make connection by tcp & returns its result in Stdout
func TCPConnect() {
	conn, err := net.Dial("tcp", "furuhama.github.io:80")
	if err != nil {
		panic(err)
	}
	conn.Write([]byte("GET / HTTP/1.0\r\nHost: furuhama.github.io\r\n\r\n"))
	io.Copy(os.Stdout, conn)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("http.ResponseWriter sample"))
}

// Handling set local server
func Handling() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}
