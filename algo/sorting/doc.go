/*
Package sort provides a collection of sorting algorithms implemented in Go.

The package includes the following functions:

 1. BubbleSort: Implements the bubble sort algorithm to sort a slice of integers.
    Bubble sort has a runtime complexity of O(n^2) in the worst case and O(n) in the best case.

 2. MergeSort: Implements the merge sort algorithm to sort a slice of integers.
    Merge sort is a divide-and-conquer algorithm with a runtime complexity of O(n * logn).

 3. Merge: Merges two sorted slices of integers into a single sorted slice.
    This function ensures that both input slices are sorted before merging them.

 4. MergeSimple: Merges two sorted slices of integers into a single sorted slice.
    This function demonstrates the simplicity and power of the standard library.

Example usage:

	package main

	import (
		"fmt"
		"github.com/eng618/go-eng/algo/sorting"
	)

	func main() {
		data := []int{5, 4, 3, 2, 1}
		sorted := sorting.BubbleSort(data)
		fmt.Println(sorted) // Output: [1 2 3 4 5]

		data = []int{5, 4, 3, 2, 1}
		sorted = sorting.MergeSort(data)
		fmt.Println(sorted) // Output: [1 2 3 4 5]

		left := []int{1, 3, 5}
		right := []int{2, 4, 6}
		merged := sorting.Merge(left, right)
		fmt.Println(merged) // Output: [1 2 3 4 5 6]

		// Example usage of Standard
		leftSimple := []int{1, 3, 5}
		rightSimple := []int{2, 4, 6}
		standard := sorting.Standard(leftSimple, rightSimple)
		fmt.Println(standard) // Output: [1 2 3 4 5 6]
	}
*/
package sorting
