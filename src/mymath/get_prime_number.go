package mymath

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
)

// GetPrimeNumber usage
// this works to get any number of prime number
// this needs the number of prime number you want
// 			& the number of algorithm for searching
// this only accept command line args (should be called from main() )
// [example] go run piyopiyo.go --alg=1 10000
func GetPrimeNumber() {
	// parsing command line args
	algno := flag.Int("alg", 0, "0: Basic algorithm, 1: Sieve of Eratosthenes")
	flag.Parse()
	args := flag.Args()
	if *algno < 0 || *algno > 1 {
		fmt.Fprintln(os.Stderr, os.ErrInvalid)
		return
	}
	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, os.ErrInvalid)
		return
	}
	max, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	if max <= 0 {
		max = 1
	}

	// searching prime number
	algoName := string("")
	prime := int64(0)
	start := time.Now() // measure time from here
	switch *algno {
	case 1:
		algoName = "Sieve of Eratosthenes algorithm"
		prime = LastPrimeE(max)
	default:
		algoName = "Basic algorithm"
		prime = LastPrimeB(max)
	}
	goal := time.Now()
	fmt.Printf("Using: %v\n", algoName)
	fmt.Printf("%vth prime number is %v\n", max, prime) // get prime number
	fmt.Printf("time: %v", goal.Sub(start))
}

// Basic algorithm
func LastPrimeB(max int64) int64 {
	count := int64(0)

	for n := int64(2); ; n++ {
		flag := true
		for m := int64(2); m < n; m++ {
			if (n % m) == 0 {
				flag = false
				break
			}
		}
		if flag {
			count++
			if count >= max {
				return n
			}
		}
	}
}

// Sieve of Eratosthenes algorithm
func LastPrimeE(max int64) int64 {
	if max <= 1 {
		return 2
	}
	primes := make([]int64, 1, max)     // list of prime numbers
	primes_f := make([]float64, 1, max) // list of prime numbers (cast to float)
	primes[0] = 2                       // 2 is a prime number
	primes_f[0] = 2.0                   // 2 is a prime number

	count := int64(1)
	for n := int64(3); ; n += 2 { // just checking odd numbers from 3
		flag := true
		f := float64(n)    // casting n to float
		rf := math.Sqrt(f) // get √n
		for i := 1; i < len(primes); i++ {
			if primes_f[i] > rf { // just checking n by numbers smaller than √n
				break
			} else if (n % primes[i]) == 0 {
				flag = false
				break
			}
		}
		if flag {
			count++
			if count >= max {
				return n
			}
			primes = append(primes, n)     // add prime number
			primes_f = append(primes_f, f) // add prime number
		}
	}
}
