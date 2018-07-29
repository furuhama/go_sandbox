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

	for i := 0; i < len(t); {
		r, size := utf8.DecodeRuneInString(t[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	for i, r := range t {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	u := "プログラム"
	fmt.Printf("% x\n", u)
	v := []rune(u)
	fmt.Printf("%x\n", v)
}
