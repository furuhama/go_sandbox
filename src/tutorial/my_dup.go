// this returns lines which appears twice or more in stdout
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
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
