package main

import "fmt"

// Go won't read this line

/*
Everything here won't be read either (multi-line comment)

Today: we'llñ see four variables types
int : integers
uint : nonnegative integers
bool : true or false Boolean variable
float64: decimal number

Next time:
byte : single symbol
string : contiguos collection of suymbols (words)
*/

func main() {
	fmt.Println("Let's get started!")

	var j int = 14 // j is an integer variable
	var x float64 = -2.3
	var yo string = "Hi" // default value for string is ""
	var u uint = 14
	var symbol byte = 'H'     // prints 72?
	var statement bool = true // dafaults to false

	fmt.Println(j)
	fmt.Println(x)
	fmt.Println(yo)
	fmt.Println(u)
	fmt.Println(symbol)
	fmt.Println(statement)

	//shorthand declarations
	i := -6     // automatically type int (if you want a uint, declare it)
	hi := "Yo " //Go automatically gives this type string
	k := 34     // automatically an int
	//secondStatement := true

	// we can do arithmetic on numeric variables

	fmt.Println((i + j) * 2 * k)
	fmt.Println(2*x - 3.16)

	fmt.Println(hi + yo)

	fmt.Println(k / j) // Go views as integer division ... throes away the reminder

	// if we want the actual value k/j, we use type conversions.
	fmt.Println(float64(k) / float64(j))

	// not all type conversions will work

	/*
		var ok bool = bool(0) // false ?
		fmt.Println(ok)
	*/
	var p int = -187
	var s uint = uint(p)

	fmt.Println(s)

	m := 9223372036854775807 // maximum allowed integer on memmory 2^63 - 1

	fmt.Println(m + 1)

	fmt.Println(float64(j) * x)

	/*
		// Go demands that input variables hace correct type
		w, z := DoubleAndDuplicate(m)
		fmt.Println(m)
	*/

	n := 17
	fmt.Println(addOne(n))
	fmt.Println(n)

	o := SumTwoInts(n, p)
	fmt.Println(o)
}

func addOne(n int) int { // the variable n gets created as a copy od whatever it is given.
	n = n + 1
	return n // we are returning a value
	// the new n is now destroyed at the end of function
}

//SumTwoInts takes two integers and returns their sum.
func SumTwoInts(a int, b int) int {
	return a + b
}

//DoubleAndDuplicate take one input, returns two parameters
func DoubleAndDuplicate(x float64) (float64, float64) {
	return 2.0 * x, 2.0 * x
}

//Pi doesnt take inputs, returns a float
func Pi() float64 {
	return 3.14
}
