// Package sort is a collection of sorting algorithms.
//
// See https://visualgo.net/en/sorting for a visual example of merge sort.
package sort

import "slices"

// MergeSort takes the provided data (slice of int) and applies the merge sort algorithm to sort the data.
// The runtime of merge sort is at best, at worst, and at average always O(n * logn).
func MergeSort(d []int) []int {
	// base case
	if len(d) <= 1 {
		return d
	}

	// split data into 2 halves
	middle := len(d) / 2
	left := d[:middle]
	right := d[middle:]

	// recursively merge sorted sides
	return Merge(MergeSort(left), MergeSort(right))
}

// Merge takes two slices of integers, l and r, and returns a new slice containing
// all elements from both input slices, sorted in ascending order. The function
// ensures that both input slices are sorted before merging them. It then iterates
// through both slices, adding the smaller element from either slice to the result
// slice until all elements from both input slices have been added.
//
// Parameters:
//   - l: A slice of integers.
//   - r: A slice of integers.
//
// Returns:
//
//	A new slice of integers containing all elements from l and r, sorted in ascending order.
func Merge(l, r []int) []int {
	result := make([]int, len(l)+len(r))

	// Ensure both sides are sorted.
	slices.Sort(l)
	slices.Sort(r)

	// Add items to result until either side is empty.
	li, ri, i := 0, 0, 0
	for li < len(l) || ri < len(r) {
		if li < len(l) && (ri == len(r) || l[li] < r[ri]) {
			result[i] = l[li]
			li++
		} else {
			result[i] = r[ri]
			ri++
		}
		i++
	}

	return result
}
