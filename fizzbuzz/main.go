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
	if f && b {
		return "FizzBuzz"
	}
	if f {
		return "Fizz"
	}
	if b {
		return "Buzz"
	}
	return ""
}

func main() {
	for i := 0; i <= 100; i++ {
		fmt.Printf("%d: %s\n", i, fizzerBuzzer(i))
	}
	os.Exit(0)
}
