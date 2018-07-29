package ch3

import (
	"fmt"
)

// BasicType of Golang
func BasicType() {
	// basic integer types
	var u uint8 = 255
	fmt.Println(u, u+1, u*u)

	var i int8 = 127
	fmt.Println(i, i+1, i*i)

	// bit calculation
	var b1 uint8 = 1<<1 | 1<<5
	var b2 uint8 = 1<<1 | 1<<2

	fmt.Printf("%08b\n", b1)
	fmt.Printf("%08b\n", b2)

	fmt.Printf("%08b\n", b1&b2)
	fmt.Printf("%08b\n", b1|b2)
	fmt.Printf("%08b\n", b1^b2)
	fmt.Printf("%08b\n", b1&^b2)

	// if you use uint for access to array,
	// after i becomes 0, i-- will be 2^64 - 1 (it depends on i's bit size)
	// therefore, in this case, you should use not unsigned int (like int64)
	medals := []string{"gold", "silver", "bronze"}
	for i := len(medals) - 1; i >= 0; i-- {
		fmt.Println(medals[i])
	}
}
