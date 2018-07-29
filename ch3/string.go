package ch3

import (
	"fmt"
	"strings"
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

	fmt.Println(string(65))      // => A
	fmt.Println(string(0x4eac))  // => 京
	fmt.Println(string(1234567)) // => �

	fmt.Println(basenameImproved("a/b/c.go")) // => c
	fmt.Println(basenameImproved("c.d.go"))   // => c.d
	fmt.Println(basenameImproved("abc"))      // => abc
}

func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}

	return s
}

func basenameImproved(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}

	return s
}
