package main

import (
	"fmt"
	"systems"
)

func main() {
	fmt.Println("TCPSocketServer start")
	systems.TCPSocketServerChunk()
	fmt.Println("TCPSocketServer end")
}
