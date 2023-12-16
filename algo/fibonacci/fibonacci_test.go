package fibonacci_test

import (
	"fmt"

	"github.com/eng618/go-eng/algo/fibonacci"
)

func ExampleBasic() {
	// Fib example
	fmt.Println(fibonacci.Basic(10))

	// Output:
	// 55
}

func ExampleBottomUp() {
	// FibonacciBU example
	fmt.Println(fibonacci.BottomUp(10))

	// Output:
	// 55
}

func ExampleDynamic() {
	// Fibonacci example
	fibO := fibonacci.Dynamic()
	fmt.Println(fibO(10))

	// Output:
	// 55
}
