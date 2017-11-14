// use Sieve of Eratosthenes
// refactor prime_num_2.go to delete unnecessary calculation
package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	primes := make([]int64, 1)     // list of prime number
	primes_f := make([]float64, 1) // list of prime number (cast to float)
	primes[0] = 2                  // 2 is a prime number
	primes_f[0] = 2.0              // 2 is a prime number (float)
	var max int64 = 10000

	start := time.Now() // measure time from here
	var n int64 = 3
	for n = 3; n < max; n += 2 { // check only odd number from 3
		flag := true
		f := float64(n)                    // casting n to float
		rf := math.Sqrt(f)                 // get √n
		for i := 1; i < len(primes); i++ { // check by already known prime number greater than 2
			if primes_f[i] > rf { // just check numbers smaller than √n
				break
			} else if (n % primes[i]) == 0 { // except numbers divisible by prime numbers
				flag = false
				break
			}
		}
		if flag {
			primes = append(primes, n)     // add prime number
			primes_f = append(primes_f, f) // add prime number(float)
		}
	}
	goal := time.Now()
	fmt.Printf("prime numbers under %v are %v\n", max, primes)
	fmt.Printf("time: %v", goal.Sub(start))
}
