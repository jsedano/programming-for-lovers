package main

func main() {

}

func simpleFunction(x, y int) int {
	if x == y { // == means testing wether two expression are equal
		return 0
	} else if x > y {
		return 1
	} else { // we can infer x < y
		return -1
	}
}

func sameSign(x, y int) bool {
	if x*y >= 0 { // >= is "greater tha or equal"
		return true
	}
	// if we make it here, we know that x * y >= 0
	return false
}

// PositiveDifference when we declare a variable, it has a "scope" (lifetime) that exists inside
// of the "block" in which it lives.
func PositiveDifference(a, b int) int {
	var c int
	if a == b {
		c = 0
	} else if a > b {
		c = a - b
	} else { // a < b
		c = b - a
	}
	return c
}

// Other conditions
/*
>, <, >=. <=, ==

"not equal to" !=
"not" !x (this evaluates to true when x is false and false when x is true)

*/
