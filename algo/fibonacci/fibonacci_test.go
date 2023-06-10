package fibonacci_test

import (
	"fmt"

	"github.com/eng618/go-eng/algo/fibonacci"
)

func Example() {
	// Fibonacci example
	fibO := fibonacci.Dynamic()
	fmt.Println(fibO(10))

	// FibonacciBU example
	fmt.Println(fibonacci.BottomUp(10))

	// Fib example
	fmt.Println(fibonacci.Basic(10))

	// Output:
	// 55
	// 55
	// 55
}
