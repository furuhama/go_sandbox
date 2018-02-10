package main

import (
	"hello"
	"systems"
)

func main() {
	hello.Hello()

	systems.StdOut()
	systems.Buffer()
	// systems.TCPConnect()
	// systems.Handling()
	systems.FlushBuf()
}
