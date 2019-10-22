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
