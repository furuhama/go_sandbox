// Package systems is for system layor program
package systems

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
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
		go processSession(conn)
	}
}

func processSession(conn net.Conn) {
	fmt.Printf("Accept %v\n", conn.RemoteAddr())
	defer conn.Close()

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

		// write response
		// setting for HTTP/1.1 & ContentLength
		response := http.Response{
			StatusCode: 200,
			ProtoMajor: 1,
			ProtoMinor: 1,
			Header:     make(http.Header),
		}

		if isGZipAcceptable(request) {
			content := "Hello, World (gzipped)\n"
			// transfer contents as gzipped data
			var buffer bytes.Buffer
			writer := gzip.NewWriter(&buffer)
			io.WriteString(writer, content)
			writer.Close()
			response.Body = ioutil.NopCloser(&buffer)
			response.ContentLength = int64(buffer.Len())
			response.Header.Set("Content-Encoding", "gzip")
		} else {
			content := "Hello, World\n"
			response.Body = ioutil.NopCloser(strings.NewReader(content))
			response.ContentLength = int64(len(content))
		}

		response.Write(conn)
	}
}

// http://www.aozora.gr.jp/cards/000121/card628.html
var longContents = []string{
	"これは、私わたしが小さいときに、村の茂平もへいというおじいさんからきいたお話です。",
	"むかしは、私たちの村のちかくの、中山なかやまというところに小さなお城があって、",
	"中山さまというおとのさまが、おられたそうです。",
	"その中山から、少しはなれた山の中に、「ごん狐ぎつね」という狐がいました。",
	"ごんは、一人ひとりぼっちの小狐で、しだの一ぱいしげった森の中に穴をほって住んでいました。",
	"そして、夜でも昼でも、あたりの村へ出てきて、いたずらばかりしました。",
}

func processSessionChunk(conn net.Conn) {
	fmt.Printf("Accept %v\n", conn.RemoteAddr())
	defer conn.Close()
	for {
		// read request
		request, err := http.ReadRequest(bufio.NewReader(conn))
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		dump, err := httputil.DumpRequest(request, true)
		if err != nil {
			panic(err)
		}

		fmt.Println(string(dump))

		// write response
		fmt.Fprintf(conn, strings.Join([]string{
			"HTTP/1.1 200 OK",
			"Content-Type: text/plain",
			"Transfer-Encoding: chunked",
			"", "",
		}, "\r\n"))

		for _, content := range longContents {
			bytes := []byte(content)
			fmt.Fprintf(conn, "%x\r\n%s\r\n", len(bytes), content)
		}
		fmt.Fprintf(conn, "0\r\n\r\n")
	}
}

func isGZipAcceptable(request *http.Request) bool {
	return strings.Index(strings.Join(request.Header["Accept-Encoding"], ","), "gzip") != -1
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
