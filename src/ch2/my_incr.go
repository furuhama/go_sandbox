package ch2

func MyIncr(x *int) int {
	*x++
	return *x
}
