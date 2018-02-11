// Package ch2 is chapter 2 of "The Go Programming Language"
package ch2

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
	// Array pc is below
	// [0 1 1 2 1 2 2 3 1 2 2 3 2 3 3 4]
	// [1 2 2 3 2 3 3 4 2 3 3 4 3 4 4 5]
	// [1 2 2 3 2 3 3 4 2 3 3 4 3 4 4 5]
	// [2 3 3 4 3 4 4 5 3 4 4 5 4 5 5 6]
	// [1 2 2 3 2 3 3 4 2 3 3 4 3 4 4 5]
	// [2 3 3 4 3 4 4 5 3 4 4 5 4 5 5 6]
	// [2 3 3 4 3 4 4 5 3 4 4 5 4 5 5 6]
	// [3 4 4 5 4 5 5 6 4 5 5 6 5 6 6 7]
	// [1 2 2 3 2 3 3 4 2 3 3 4 3 4 4 5]
	// [2 3 3 4 3 4 4 5 3 4 4 5 4 5 5 6]
	// [2 3 3 4 3 4 4 5 3 4 4 5 4 5 5 6]
	// [3 4 4 5 4 5 5 6 4 5 5 6 5 6 6 7]
	// [2 3 3 4 3 4 4 5 3 4 4 5 4 5 5 6]
	// [3 4 4 5 4 5 5 6 4 5 5 6 5 6 6 7]
	// [3 4 4 5 4 5 5 6 4 5 5 6 5 6 6 7]
	// [4 5 5 6 5 6 6 7 5 6 6 7 6 7 7 8]
}

// PopCount returns population count
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
