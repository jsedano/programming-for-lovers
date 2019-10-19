package main

import "fmt"

func main() {
	fmt.Println("Loops.")

	n := 5
	m := factorial(n)

	fmt.Println(m)

	//fmt.Println(factorial(-100))
	fmt.Println(sumFirstIntegers(10))

	fmt.Println(sumEven(5))

	// closing lesson: infinite loop in disguise
	/*
		var i uint = 10

		for ; i >= 0; i-- {
			fmt.Println(i)
		}
	*/

}

// first, a factorial function

func factorial(n int) int {
	// let's handle negative input
	if n < 0 {
		// panic() will print an error message and inmediately end program.
		panic("ERROR: negative input given to factorial")
	}
	p := 1
	i := 1
	//Go doesn't have a while keyword. They use "for"
	for i <= n {
		p = p * i
		i++
	}

	return p
}

func anotherFactorial(n int) int {
	if n < 0 {
		// panic() will print an error message and inmediately end program.
		panic("ERROR: negative input given to factorial")
	}
	p := 1

	// for every integer i between 1 and n, P = p*i
	for i := 1; i <= n; i++ {
		p *= i
	}
	return p
}

// Exercise: write a function in Go using a while loop that takes an integer n and returns the sum of the first n positive integers

func sumFirstIntegers(n int) (r int) {
	if n <= 0 {
		panic("Negative number")
	}
	for i := 1; i <= n; i++ {
		r += i
	}
	return //naked return
}

//Excersice: Write a function sumEven that sums all even number up to and possible including n

func sumEven(n int) (r int) {
	for i := 2; i <= n; i += 2 {
		r += 2
	}
	return
}
