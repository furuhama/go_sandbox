// Package mymath is original implementation of math functions
package mymath

import (
	"fmt"
	"time"
)

// PrimeNum returns a prime number
func PrimeNum() {
	max := 10000 // set max number
	primeList := make([]int64, 0)

	start := time.Now() // measure time from here
	for n := 2; n <= max; n++ {
		flag := true
		for m := 2; m < n; m++ {
			if (n % m) == 0 {
				flag = false
				break
			}
		}
		if flag {
			primeList = append(primeList, int64(n))
		}
	}
	goal := time.Now() // measure time end

	fmt.Printf("Prime numbers under %v are %v\n", max, primeList)
	fmt.Printf("time: %v", goal.Sub(start))
}
