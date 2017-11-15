// use Sieve of Eratosthenes
package mymath

import (
	"fmt"
	"time"
)

func PrimeNum2() {
	max := 10000                      // set max number
	primes := make([]int64, 0)        // list of prime numbers
	isNotPrime := make([]bool, max+1) // initialize as false(= prime number)
	isNotPrime[0] = true              // 0 is not prime number
	isNotPrime[1] = true              // 1 is not prime number

	start := time.Now() // measure time from here
	n := 2
	for ; n*n <= max; n++ { // put into sieve at most sqrl(max)
		if !isNotPrime[n] {
			primes = append(primes, int64(n)) // add prime number
			for m := 2; n*m <= max; m++ {
				isNotPrime[n*m] = true // set not prime number
			}
		}
	}
	for ; n <= max; n++ {
		if !isNotPrime[n] {
			primes = append(primes, int64(n)) // add prime number
		}
	}
	goal := time.Now() // measure time end
	fmt.Printf("prime numbers under %v are %v\n", max, primes)
	fmt.Printf("time: %v", goal.Sub(start))
}
