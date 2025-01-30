package fibonacci_test

import (
	"fmt"
	"testing"

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

func TestBasic(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{10, 55},
		{20, 6765},
	}

	for _, test := range tests {
		if result := fibonacci.Basic(test.input); result != test.expected {
			t.Errorf("Basic(%d) = %d; want %d", test.input, result, test.expected)
		}
	}
}

func TestBottomUp(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{10, 55},
		{20, 6765},
	}

	for _, test := range tests {
		if result := fibonacci.BottomUp(test.input); result != test.expected {
			t.Errorf("BottomUp(%d) = %d; want %d", test.input, result, test.expected)
		}
	}
}

func TestDynamic(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{10, 55},
		{20, 6765},
	}

	fibO := fibonacci.Dynamic()
	for _, test := range tests {
		if result := fibO(test.input); result != test.expected {
			t.Errorf("Dynamic(%d) = %d; want %d", test.input, result, test.expected)
		}
	}
}

func BenchmarkBasic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacci.Basic(20)
	}
}

func BenchmarkBottomUp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacci.BottomUp(20)
	}
}

func BenchmarkDynamic(b *testing.B) {
	fibO := fibonacci.Dynamic()
	for i := 0; i < b.N; i++ {
		fibO(20)
	}
}
