// Package tutorial is for the chapter 1 of "The Go Programming Language"
package ch1

import (
	"bufio"
	"fmt"
	"os"
)

// MyDup returns lines which appears twice or more in stdout
func MyDup() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	// [caution] ignore the possibility of error from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("\n%d\t%s", n, line)
		}
	}
}
