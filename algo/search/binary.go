package search

// BinaryLoop is an algorithm that uses the binary search method to search a given slice.
// We are expecting a sorted slice, otherwise the result will not be accurate.
//
// https://goplay.space/#TQBxLBIrt-w
//
// https://play.golang.org/p/TQBxLBIrt-w
func BinaryLoop(xi []int, target int) (index int, ok bool) {
	min, max := 0, len(xi)-1

	for min <= max {
		guess := (min + max) / 2

		if xi[guess] == target {
			return guess, true
		} else if xi[guess] < target {
			min = guess + 1
		} else {
			max = guess - 1
		}
	}
	return -1, false
}

// BinaryRecursion uses the recursion method to call itself, until it
// determines if the target is present in the slice (which must be sorted).
func BinaryRecursion(xi []int, target int) bool {
	if len(xi) == 0 {
		return false
	}
	if len(xi) == 1 {
		return xi[0] == target
	}
	middle := len(xi) / 2
	if xi[middle] == target {
		return true
	} else if xi[middle] < target {
		return BinaryRecursion(xi[middle+1:], target)
	} else {
		return BinaryRecursion(xi[:middle], target)
	}
}
