package tcpserver

import (
	"fmt"
	"net"
)

// ListenOnce starts server which listens port 8080
// and which responses message just once.
// (quickly close connection just after send response)
func ListenOnce() {
	fmt.Println("Start simple TCP server")

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	conn, err := ln.Accept()
	if err != nil {
		panic(err)
	}

	defer func() {
		conn.Close()
		fmt.Println("Connection closed")
	}()

	_, err = conn.Write([]byte("test"))
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully send response")
}
