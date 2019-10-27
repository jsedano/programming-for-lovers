package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println(permutation1(10, 4))
	fmt.Println(combination1(10, 4))

	fmt.Println(permutation2(1000, 2).Int64())
	fmt.Println(combination2(1000, 2))
	fmt.Println(combination2(1000, 998))
	fmt.Println(factorialArray(11))
	fmt.Println(fibonacciArray(13))
	fmt.Println(minArray([]int{-1, 3, -8, -3, -5, 7, -2, 9, -4, 0}))
	fmt.Println(gcdArray([]int{96, 12, 6, 48, 24}))

	fmt.Println("perfect numbers:")

	for i := 1; i < 10000; i++ {
		if isPerfect(i) {
			fmt.Println(i)
		}
	}
	fmt.Println("is 33550336 perfect: ", isPerfect(33550336))

}

//Exercise 1 write and implement Combination and Permutacion

func permutation1(n, k int) int {
	return factorial(n) / factorial(n-k)
}

func combination1(n, k int) int {
	return permutation1(n, k) / factorial(k)
}

//Exercise 2 modify Combination and Permutation so they are able to compute
//permutation(1000, 2))
//combination(1000, 2))
//combination(1000, 998))

func permutation2(n, k int64) *big.Int {
	x := factorial1(*big.NewInt(n))
	y := factorial1(*big.NewInt(n).Sub(big.NewInt(n), big.NewInt(k)))
	z := *x.Div(&x, &y)
	return &z
}

func combination2(n, k int64) int64 {
	x := permutation2(n, k)
	y := factorial1(*big.NewInt(k))
	z := *x.Div(x, &y)
	return z.Int64()
}

func factorial(n int) int {
	if n < 0 {
		panic("ERROR: negative input given to factorial")
	}
	p := 1
	for i := 1; i <= n; i++ {
		p *= i
	}
	return p
}

func factorial1(n big.Int) big.Int {
	if n.Cmp(big.NewInt(0)) == -1 {
		panic("ERROR: negative input given to factorial")
	}
	var p = big.NewInt(1)
	for i := big.NewInt(1); i.Cmp(&n) <= 0; i.Add(i, big.NewInt(1)) {
		if n.Cmp(big.NewInt(0)) == 0 {

			panic("ERROR: p == 0")

		}
		p.Mul(p, i)
	}
	return *p
}

// Working with arrays

// implement a function factorialArray that takes an integer n and returns a slice of lenght n + 1 whose k-th element is equal to k!

func factorialArray(n int) []int {
	r := make([]int, n+1)
	r[0] = 1
	for i := 1; i < len(r); i++ {
		r[i] = i * r[i-1]
	}
	return r
}

// implement a function fibonacciArray that takes an integer n and returns a slice of lenght n + 1 whose k-th element is the k-th Fibonacci number

func fibonacciArray(n int) []int {
	r := make([]int, n)
	r[0] = 0
	r[1] = 1
	for i := 2; i < len(r); i++ {
		r[i] = r[i-1] + r[i-2]
	}
	return r
}

// write and implement a function minArray that takes an array of integers as input and returns the minimum of all these integers

func minArray(arr []int) int {
	min := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[min] {
			min = i
		}
	}
	return arr[min]
}

// write and implement a function gcdArray that takes an array of integers as input
//and generalizes the idea in trivial gcd to return the greatest common divisor of all the integers in the array

func gcdArray(arr []int) int {
	d := 1
	m := minArray(arr)
	for p := 1; p <= m; p++ {
		// remainder operation Reminder(n, k) is n%k (e,g,, 14%3 = 1)
		tmpd := d
		for _, v := range arr {
			if v%p == 0 {
				tmpd = p
			} else {
				tmpd = -1
				break
			}
		}
		if tmpd != -1 {
			d = tmpd
		}
	}
	return d
}

// write isPerfect that takes an integer n as input and returns true if n is perfect and false otherwise

func isPerfect(n int) bool {
	divisors := make([]int, 0)
	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
		}
	}
	sum := 0
	for _, v := range divisors {
		sum += v
	}
	return sum == n
}
