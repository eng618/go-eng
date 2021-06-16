package search

// Linear is a simple search algorithm to verify if a target element is present
// in the provided slice. Unlike the binary search, this does not require the
// input to be sorted prior to running.
//
// On average the complexity of this algorithm is O(n).
func Linear(xi []int, target int) bool {
	for _, v := range xi {
		if v == target {
			return true
		}
	}

	return false
}
