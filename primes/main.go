package main

import (
	"fmt"
	"log"
	"math"
	"time"
)

func main() {
	fmt.Println("Primes and arrays.")

	// arrays in go have a fixed, constant size
	var list [6]int
	list[0] = -8

	i := 3
	list[2*i-4] = 17
	list[len(list)-1] = 43
	// compiler errors (index out of bounds)
	//list[len(list)] = 4
	//list[-4] = 0
	fmt.Println(list)

	start := time.Now()
	TrivialPrimeFinder(1000000)
	elapsed := time.Since(start)
	log.Printf("TrivialPrimeFinder %s", elapsed)

	start2 := time.Now()
	SieveOfEratosthenes(100000)
	elapsed2 := time.Since(start2)
	log.Printf("TrivialPrimeFinder %s", elapsed2)
}

// remember that arrays have constant sizer in Go.

// we use something called a "slice" to reprensent variable sizes in Go.
// even if you're plugging in a variable into the length of the array and will never change it

//TrivialPrimeFinder takes an integer n and produces a boolean array of length n+1 where primeArray[p] is true if p is prime (amd false otherwise)
func TrivialPrimeFinder(n int) []bool {
	var primeArray []bool // type is "slice of booleans"
	// we need an initial length.
	// slices start as nil and need to be made
	primeArray = make([]bool, n+1)
	// in general, you'll just use primeArray := make([]bool, n+1)
	for p := 2; p <= n; p++ {
		if IsPrime(p) {
			primeArray[p] = true
		}
	}

	return primeArray
}

//IsPrime checks if number is prime
func IsPrime(p int) bool {
	for k := 2; float64(k) <= math.Sqrt(float64(p)); k++ {
		if p%k == 0 { // p is not prime
			return false
		}
	}
	//if we survice testing all these factors, p is prime
	return true
}

//SievesOfEratosthenes takes an integer n and returns a slice of n+1 booleans prime array where primeArray[p] is true if p is prime
func SieveOfEratosthenes(n int) []bool {
	primeArray := make([]bool, n+1)

	// set everyuthing to prime other than 0 and 1
	for k := 2; k <= n; k++ {
		primeArray[k] = true
	}

	// now range over primeArray, and cross off multiples of first prime we see and iterate this process
	for p := 2; float64(p) <= math.Sqrt(float64(n)); p++ {
		if primeArray[p] == true {
			primeArray = CrossOffMultiples(primeArray, p)
		}
	}

	return primeArray
}

//CrossOffMultiples takes slice and integer p, it crosses off multiples of p, meaning turning there multiple to false in the slice
//it returns the slice obtained as result
func CrossOffMultiples(primeArray []bool, p int) []bool {
	n := len(primeArray) - 1
	for k := 2 * p; k <= n; k += p {
		// all there multiples shoud be made composite
		primeArray[k] = false
	}
	return primeArray
}
