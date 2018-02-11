// Package mymath is original implementation of math functions
package mymath

import (
	"fmt"
	"mymath"
)

// PrintSqrt prints sqrt of arg in stdout
func PrintSqrt(x float64) {
	fmt.Printf("Hello, world. Sqrt(%v) = %v\n", x, mymath.Sqrt(x))
}
