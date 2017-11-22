package mymath

// MyFib returns fibonacci number
func MyFib(n int) int64 {
	var x, y int64 = 1, 0
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
