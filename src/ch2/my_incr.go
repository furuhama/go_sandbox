// Package ch2 is chapter 2 of "The Go Programming Language"
package ch2

func MyIncr(x *int) int {
	*x++
	return *x
}
