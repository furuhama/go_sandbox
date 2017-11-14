package main

import(
	"fmt"
	"time"
)

func main() {
	max := 10000 // set max number
	prime_list := make([]int64, 0)

	start := time.Now() // measure time from here
	for n := 2; n <= max; n++ {
		flag := true
		for m := 2; m < n; m++ {
			if (n % m) == 0 {
				flag = false
				break
			}
		}
		if flag {
			prime_list = append(prime_list, int64(n))
		}
	}
	goal := time.Now() // measure time end

	fmt.Printf("Prime numbers under %v is %v\n", max, prime_list)
	fmt.Printf("time: %v", goal.Sub(start))
}
