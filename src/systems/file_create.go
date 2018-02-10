// Package systems is for system layor program
package systems

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
)

// CreateText is making txt file & write byte data in it
func CreateText() {
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	file.Write([]byte("os.File examples"))
	file.Close()
}

// StdOut send Stdout to example texts
func StdOut() {
	os.Stdout.Write([]byte("os.Stdout example\n"))
}

// Buffer put byte data into buffer, and send it to Stdout
func Buffer() {
	var buffer bytes.Buffer
	buffer.Write([]byte("bytes.Buffer example\n"))
	fmt.Println(buffer.String())
}

// TCPConnect make connection by tcp & returns its result in Stdout
func TCPConnect() {
	conn, err := net.Dial("tcp", "furuhama.github.io:80")
	if err != nil {
		panic(err)
	}
	conn.Write([]byte("GET / HTTP/1.0\r\nHost: furuhama.github.io\r\n\r\n"))
	io.Copy(os.Stdout, conn)
}
