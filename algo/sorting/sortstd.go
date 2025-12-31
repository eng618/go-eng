package sorting

import "slices"

// Standard merges two sorted slices into a single sorted slice.
// It takes two slices of integers, l and r, and returns a new slice
// containing all elements from both input slices, sorted in ascending order.
// This function demonstrates the simplicity and power of the standard library.
//
// Parameters:
//   - l: A sorted slice of integers.
//   - r: A sorted slice of integers.
//
// Returns:
//
//	A new sorted slice containing all elements from both input slices.
func Standard(l, r []int) []int {
	result := append(l, r...)
	slices.Sort(result)
	return result
}
