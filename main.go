package main

import (
	"fmt"
	"systems"
)

func main() {
	fmt.Println("TCPSocketClient start")
	systems.TCPSocketClientPipeline()
	fmt.Println("TCPSocketClient end")
}
