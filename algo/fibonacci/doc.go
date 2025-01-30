/*
Package fibonacci provides functions to calculate Fibonacci numbers using different approaches.

The Fibonacci sequence is a series of numbers where each number is the sum of the two preceding ones,
usually starting with 0 and 1. The sequence goes as follows:

	0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, ...

This package includes the following functions:

1. Basic: A simple recursive implementation with exponential time complexity.
2. BottomUp: An iterative implementation with linear time complexity.
3. Dynamic: A memoized recursive implementation with linear time complexity.

Example usage:

	package main

	import (
		"fmt"
		"github.com/eng618/go-eng/algo/fibonacci"
	)

	func main() {
		fmt.Println(fibonacci.Basic(10))    // Output: 55
		fmt.Println(fibonacci.BottomUp(10)) // Output: 55
		fib := fibonacci.Dynamic()
		fmt.Println(fib(10))                // Output: 55
	}
*/
package fibonacci
