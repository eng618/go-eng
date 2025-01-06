package search

// BinaryLoop is an algorithm that uses the binary search method to search a given slice.
// We are expecting a sorted slice, otherwise the result will not be accurate.
//
// https://goplay.space/#TQBxLBIrt-w
//
// https://play.golang.org/p/TQBxLBIrt-w
func BinaryLoop(xi []int, target int) (index int, ok bool) {
	minimum, maximum := 0, len(xi)-1

	for minimum <= maximum {
		guess := (minimum + maximum) / 2

		switch {
		case xi[guess] == target:
			return guess, true
		case xi[guess] < target:
			minimum = guess + 1
		default:
			maximum = guess - 1
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

	switch {
	case xi[middle] == target:
		return true
	case xi[middle] < target:
		return BinaryRecursion(xi[middle+1:], target)
	default:
		return BinaryRecursion(xi[:middle], target)
	}
}
