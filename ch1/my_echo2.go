// Package ch1 is for the chapter 1 of "The Go Programming Language"
package ch1

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// MyEcho2 behaves like "echo" linux command in two different ways
// then, returns the time of each way
func MyEcho2() {
	fmt.Println("Pattern1 start...")
	start1 := time.Now()
	// Patten 1
	// using range
	s, sep := "", ""
	for _, arg := range os.Args[1:] { // range generates index, value
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	goal1 := time.Now()

	fmt.Println("Pattern2 start...")
	start2 := time.Now()
	// Pattern 2
	// this returns process name & result
	// fmt.Printf("Process name:\n%v\n\n", os.Args[0])
	// fmt.Printf("Result:\n")
	fmt.Printf(strings.Join(os.Args[1:], " "))
	goal2 := time.Now()

	// Pattern 3
	// index, value
	//for key, val := range os.Args[1:] {
	//	fmt.Printf("%v: %v", key, val)
	//}

	// measure time
	fmt.Printf("\nPattern1:\n%v\nPattern2:\n%v", goal1.Sub(start1), goal2.Sub(start2))
}
