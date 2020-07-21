package main

import (
	"fmt"
	"os"
)

// fizzerBuzzer returns "Fizz" if the passed integer is divisible by 3,
// "Buzz" if it is divisible by 5, "FizzBuzz" if it is divisible by both, and
// an empty string if divisible by neither.
func fizzerBuzzer(n int) string {
	f := (n%3 == 0)
	b := (n%5 == 0)

	switch {
	case f && b:
		return "FizzBuzz"
	case f:
		return "Fizz"
	case b:
		return "Buzz"
	default:
		return ""
	}
}

// Time complexity: O(n) linear because it prints a line for each number.
// To perform the calculation for one n is O(1) constant.
// Memory complexity: O(1)
func main() {
	for i := 0; i <= 100; i++ {
		fmt.Printf("%d: %s\n", i, fizzerBuzzer(i))
	}
	os.Exit(0)
}
