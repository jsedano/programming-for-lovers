package main

import (
	"fmt"
	"log"
	"time"
)

/*
TrivialGDC(a,b)
	d <- 1
	m <- Min2(a, b)
	for every integer p from 1 to m
		if p is a divisor of a and b
			d <- p
	return d
*/

func main() {
	x := 378
	y := 273

	start := time.Now()
	fmt.Println(trivialGDC(x, y))
	elapsed := time.Since(start)
	log.Printf("Trivial GCD took %s", elapsed)

	start2 := time.Now()
	fmt.Println(euclidGCD(x, y))
	elapsed2 := time.Since(start2)
	log.Printf("Euclid GCD took %s", elapsed2)
}

func trivialGDC(a, b int) int {
	d := 1
	m := min2(a, b)
	for p := 1; p <= m; p++ {
		// remainder operation Reminder(n, k) is n%k (e,g,, 14%3 = 1)
		if a%p == 0 && b%p == 0 {
			d = p
		}
	}
	return d
}

/*
EuclidGCD(a, b)
	while a not equal to b
		if a > b
			a = a - b
		else // b < a
			b = b - a
	return a // or b
*/

func euclidGCD(a, b int) int {
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}
	return a
}

/*
	&& "AND" operator, a && f is true if both are true
	|| "OR" operator, a || f is true if at least one is true
*/

func min2(a, b int) int {
	if a < b {
		return a
	}
	return b
}
