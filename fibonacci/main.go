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
	ppn := 0 // previous previous number
	pn := 0  // previous number
	cn := 1  // current number
	if n == 0 {
		return 0
	}
	for i := 1; i < n; i++ {
		ppn = pn
		pn = cn
		cn = ppn + pn
	}
	return cn
}

func main() {
	fmt.Printf("Fib(n) = recursive = iterative\n")
	for i := 0; i <= 20; i++ {
		fmt.Printf("Fib(%d) = %d = %d\n", i, fibR(i), fibI(i))
	}
}
