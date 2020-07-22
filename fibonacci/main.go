package main

import "fmt"

// f(0) == 0
// f(1) == 1
// f(n) == f(n-1) + f(n-2)

// FibR recursively calculates and returns the nth Fibonacci number.
// Time complexity: O(2^n) exponential
// Memory complexity: O(n) linear
func fibR(n int) int {
	if n <= 1 {
		return n
	}
	return fibR(n-1) + fibR(n-2)
}

// FibI iteratively calculates and returns the nth Fibonacci number.
// Time complexity: O(n) linear
// Memory complexity: O(1) constant
func fibI(n int) int {
	// error condition (could also use uint in signature)
	if n < 0 {
		return -1
	}

	// base cases
	if n < 2 {
		return n
	}

	ppn := 0 // previous previous number
	pn := 1  // previous number
	cn := 0  // current number

	// note the unusual `<=` here. I think this makes more sense than starting
	// `i` at `1`.
	for i := 2; i <= n; i++ {
		// calculate the new current number
		cn = ppn + pn
		// prep for next ireration by moving values
		ppn = pn
		pn = cn
	}

	return cn
}

func main() {
	fmt.Printf("Fib(n) = recursive = iterative\n")
	for i := 0; i <= 20; i++ {
		fmt.Printf("Fib(%d) = %d = %d\n", i, fibR(i), fibI(i))
	}
}
