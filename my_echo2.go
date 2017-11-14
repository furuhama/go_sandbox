package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] { // range generates index, value
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
