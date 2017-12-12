// Package mathapp is my math app
package mathapp

import (
	"fmt"
	"mymath"
)

// PrintSqrt prints sqrt of arg in stdout
func PrintSqrt(x float64) {
	fmt.Printf("Hello, world. Sqrt(%v) = %v\n", x, mymath.Sqrt(x))
}
