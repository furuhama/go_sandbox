// Package mymath is original implementation of math functions
package mymath

import (
	"fmt"
)

// Fizzbuzz is fizzbuzz program from 1 to 30
func Fizzbuzz() {
	for i := 1; i <= 30; i++ {
		if i%15 == 0 {
			fmt.Println("fizzbuzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}
