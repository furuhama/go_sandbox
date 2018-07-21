// Package mymath is original implementation of math functions
// Source code of $GOPATH/src/mymath/sqrt.go
package mymath

import (
	"fmt"
)

// Sqrt returns the square root
func Sqrt(x float64) float64 {
	z := 0.0
	for i := 0; i < 1000; i++ {
		z -= (z*z - x) / (2 * x)
	}
	return z
}

// PrintSqrt prints sqrt of arg in stdout
func PrintSqrt(x float64) {
	fmt.Printf("Hello, world. Sqrt(%v) = %v\n", x, Sqrt(x))
}
