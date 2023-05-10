package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	prev, curr := 0, 1

	return func() int {
		result := prev
		prev, curr = curr, prev+curr
		return result
	}
}

func main() {
	// Create a Fibonacci closure
	f := fibonacci()

	// Generate and print the first 10 Fibonacci numbers
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}