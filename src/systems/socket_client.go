// Package systems is for system layor program
package systems

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
)

// TCPSocketClient sends http request
// to be compatible with keep-alive
func TCPSocketClient() {
	sendMessages := []string{
		"PIYO",
		"FUNCTION",
		"PLUSPLUS",
	}
	current := 0
	var conn net.Conn
	for {
		var err error
		// start Dial when haven't set connection, or retry because of an error
		if conn == nil {
			conn, err = net.Dial("tcp", "localhost:8888")
			if err != nil {
				panic(err)
			}
			fmt.Printf("Access: %d\n", current)
		}

		// define request to send literals as POST method
		request, err := http.NewRequest("POST",
			"http://localhost:8888",
			strings.NewReader(sendMessages[current]))
		if err != nil {
			panic(err)
		}

		request.Header.Set("Accept-Encoding", "gzip")
		err = request.Write(conn)
		if err != nil {
			panic(err)
		}

		// read response from server
		// when timeout, an error occurs written below
		response, err := http.ReadResponse(
			bufio.NewReader(conn), request)
		if err != nil {
			fmt.Println("Retry")
			conn = nil
			continue
		}

		// show results
		dump, err := httputil.DumpResponse(response, false)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dump))

		defer response.Body.Close()
		if response.Header.Get("Content-Encoding") == "gzip" {
			reader, err := gzip.NewReader(response.Body)
			if err != nil {
				panic(err)
			}
			io.Copy(os.Stdout, reader)
			reader.Close()
		} else {
			io.Copy(os.Stdout, response.Body)
		}

		// end if every content is transported
		current++
		if current == len(sendMessages) {
			break
		}
	}
}
