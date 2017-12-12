// Package mymath is original implementation of math functions
// Source code of $GOPATH/src/mymath/sqrt.go
package mymath

// Sqrt returns the square root
func Sqrt(x float64) float64 {
	z := 0.0
	for i := 0; i < 1000; i++ {
		z -= (z*z - x) / (2 * x)
	}
	return z
}
