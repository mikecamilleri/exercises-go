package main

import (
	"fmt"
	"math/rand"
	"time"
)

// boss generates work/jobs/tickets (integers 0..49 in this case) and emits
// them on workChan.
func boss(workChan chan<- int) {
	for {
		workChan <- rand.Int() % 50
	}
}

// worker does work 🛠️!
func worker(workChan <-chan int, productsChan chan<- int) {
	for job := range workChan {
		productsChan <- fibR(job)
	}
}

// FibR recursively calculates and returns the nth Fibonacci number.
// O(2^n)
func fibR(n int) int {
	if n <= 1 {
		return n
	}
	return fibR(n-1) + fibR(n-2)
}

func main() {
	rand.Seed(time.Now().Unix())

	workChan := make(chan int, 64)
	productsChan := make(chan int, 64)

	// The boss 👩‍💼 makes the work!
	// Work is distributed to several workers (fan-out pattern).
	go boss(workChan)

	// The 16 workers 👷 do the work!
	// Products are send out on a shared channel (fan-in pattern).
	for i := 0; i < 16; i++ {
		go worker(workChan, productsChan)
	}

	// an alarm sounds every 5 seconds!
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	// random alien invasion!
	// - doing this as an anonymous function for no reason
	spaceChan := make(chan string)
	go func(chan<- string) {
		for {
			time.Sleep(time.Duration((rand.Int() % 60)) * time.Second)
			spaceChan <- "🛸👾🛸👾🛸👾🛸👾🛸👾🛸"
		}
	}(spaceChan)

	for {
		select {
		case fib := <-productsChan:
			fmt.Printf("%s: 📦 %d\n", time.Now().Format(time.RFC3339), fib)
		case t := <-ticker.C:
			fmt.Printf("%s: ⏰ %v\n", time.Now().Format(time.RFC3339), t.Format(time.RFC3339))
		case hack := <-spaceChan:
			fmt.Printf("%s: 👽 %s\n", time.Now().Format(time.RFC3339), hack)
		}
	}

}
