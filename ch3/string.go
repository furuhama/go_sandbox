package ch3

import (
	"fmt"
	"unicode/utf8"
)

// String tries to use strings and slices
func String() {
	s := "hello, world"
	hello := s[:5]
	world := s[7:]

	fmt.Println(hello)
	fmt.Println(world)

	t := "hello, 世界"
	fmt.Println(len(t))                    // => 13
	fmt.Println(utf8.RuneCountInString(t)) // => 9
}
