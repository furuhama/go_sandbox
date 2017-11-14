package main

import (
	"fmt"
	"os"
)

func main() {
	// Patten 1
	//
	//s, sep := "", ""
	//for _, arg := range os.Args[1:] { // range generates index, value
	//	s += sep + arg
	//	sep = " "
	//}
	//fmt.Println(s)

	// Pattern 2
	// this returns process name & result
	fmt.Printf("Process name:\n%v\n\n", os.Args[0])
	//fmt.Printf("Result:\n%v", strings.Join(os.Args[1:], " "))
	fmt.Printf("Result: \n")

	for key, val := range os.Args[1:] {
		fmt.Printf("%v: %v", key, val)
	}
}
