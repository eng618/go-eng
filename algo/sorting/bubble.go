package sorting

// BubbleSort takes the provided data (slice of int) and applies the bubble sort algorithm to sort the data.
// The runtime of bubble sort is at best O(n) and at worst O(n^2).
func BubbleSort(d []int) []int {
	n := len(d)
	for i := range d {
		for j := range n - i - 1 {
			if d[j] > d[j+1] {
				d[j], d[j+1] = d[j+1], d[j]
			}
		}
	}
	return d
}
