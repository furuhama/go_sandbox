package myregexp

import (
	"fmt"
	"regexp"
)

// TryRegExp prints regexp results
func TryRegExp() {
	r := regexp.MustCompile(`abc`)
	fmt.Println(r.MatchString("hello"))
	fmt.Println(r.MatchString("hello abc"))

	r = regexp.MustCompile(`(a|b)b`)
	fmt.Println(r.MatchString("hello abc"))
}
