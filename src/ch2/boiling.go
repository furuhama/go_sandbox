package ch2

import "fmt"

const boilingF = 212.0

// Boiling returns boiling point as °F and °C
func Boiling() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
}
