package main

import "fmt"

func del(s []int, i int) []int {
	copy(s[i:], s[i+1:])
	s = s[:len(s)-1]
	return s
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	a = del(a, 2)
	fmt.Println(a, len(a), cap(a))
	fmt.Println(a[:cap(a)])
}
