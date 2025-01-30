// Package fibonacci allows you to calculate which fibonacci number would be at a given index. Imagine the fibonacci
// sequence is represented in the form of a slice of ints, similar to the below example.
//
//	Fibonacci Sequence
//	[]int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233...}
package fibonacci

// Dynamic utilizes memoization and dynamic programming principles,
// to calculate the fibonacci number for a given index.
func Dynamic() func(n int) int {
	var fib func(int) int

	cache := map[int]int{}

	fib = func(n int) int {
		if result, ok := cache[n]; ok {
			return result
		}

		if n < 2 {
			return n
		}

		cache[n] = fib(n-1) + fib(n-2)

		return cache[n]
	}

	return fib
}

// BottomUp is a bottom up approach to finding the number.
// Runtime complexity O(n).
func BottomUp(n int) int {
	if n == 0 {
		return 0
	}

	answer := []int{0, 1}
	for i := 2; i <= n; i++ {
		answer = append(answer, (answer[i-2] + answer[i-1]))
	}

	return answer[len(answer)-1]
}

// Basic function is simple implementation, that uses recursion.
// Runtime complexity O(n^2).
func Basic(n int) int {
	if n < 2 {
		return n
	}

	return Basic(n-1) + Basic(n-2)
}
