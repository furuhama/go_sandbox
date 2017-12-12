// Package ch2 is chapter 2 of "The Go Programming Language"
package ch2

// MyIncr receives int pointer & add 1 to its value
func MyIncr(x *int) int {
	*x++
	return *x
}
