// Package systems is for system layor program
package systems

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

// TCPSocketServer sets server up
// to be compatible with keep-alive
func TCPSocketServer() {
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server is running at localhost:8888")
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go func() {
			fmt.Printf("Accept %v\n", conn.RemoteAddr())
			for {
				// set Timeout
				conn.SetReadDeadline(time.Now().Add(5 * time.Second))
				// Read request
				request, err := http.ReadRequest(bufio.NewReader(conn))
				if err != nil {
					// End process when timeout or socket closed
					// occur error except situations written above
					neterr, ok := err.(net.Error) // downcast
					if ok && neterr.Timeout() {
						fmt.Println("Timeout")
						break
					} else if err == io.EOF {
						break
					}
					panic(err)
				}
				// display request
				dump, err := httputil.DumpRequest(request, true)
				if err != nil {
					panic(err)
				}
				fmt.Println(string(dump))
				content := "Hello, world\n"
				// write response
				// setting for HTTP/1.1 & ContentLength
				response := http.Response{
					StatusCode:    200,
					ProtoMajor:    1,
					ProtoMinor:    1,
					ContentLength: int64(len(content)),
					Body:          ioutil.NopCloser(strings.NewReader(content)),
				}
				response.Write(conn)
			}
			conn.Close()
		}()
	}
}

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
		request, err := http.NewRequest("POST", "http://localhost:8888", strings.NewReader(sendMessages[current]))
		if err != nil {
			panic(err)
		}
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
		dump, err := httputil.DumpResponse(response, true)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dump))
		// end if every content is transported
		current++
		if current == len(sendMessages) {
			break
		}
	}
}
