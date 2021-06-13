// Package sort is a collection of sortting algorithms
//
// See https://visualgo.net/en/sorting for a visual example of merge sort.
package sort

type Data []int

// MergeSort takes the provided data (slice of int) and applies the merge sort algorithm, to sort the data.
// The runtime of merge sort is at best, at worst, and at average always O(n * logn)
func MergeSort(d Data) Data {
	// base case
	if len(d) <= 1 {
		return d
	}

	// split data into 2 halves
	middle := int(len(d) / 2)
	left := d[:middle]
	right := d[middle:]

	// recursively merge sorted sides
	return Merge(MergeSort(left), MergeSort(right))
}

// Merge takes two slice of ints (assuming they are sorted) and merges them
// into a single slice. If the the inputs are sorted, the resulting merge will
// preserve the correct sorted order.
//
// If the slices are not sorted the resulting slice will have unpredictable
// results.
func Merge(l, r Data) Data {
	result := make([]int, len(l)+len(r))

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
