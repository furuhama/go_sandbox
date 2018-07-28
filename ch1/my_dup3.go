// Package tutorial is for the chapter 1 of "The Go Programming Language"
package ch1

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// MyDup3 returns lines which appears twice or more in stdout
// instead of Stdin, get file as os.Args[1:]
func MyDup3() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
