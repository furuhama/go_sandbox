// Package tutorial is for the chapter 1 of "The Go Programming Language"
package tutorial

import (
	"fmt"
	"os"
)

// MyEcho behaves like "echo" linux command
func MyEcho() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
