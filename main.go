package main

import (
	"fmt"
	"systems"
)

func main() {
	fmt.Println("TCPSocketServer start")
	systems.TCPSocketServer()
	fmt.Println("TCPSocketServer end")
}
