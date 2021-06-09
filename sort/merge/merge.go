// Package merge uses the merge sort algorithm.
// The runtime of merge sort is at best, at worst, and at average always O(n * logn)
//
// See https://visualgo.net/en/sorting for a visual example of merge sort.
package merge

type Data []int

// Sort takes the provided data (slice of int) and applies the merge sort algorithm, to sort the data.
func Sort(d Data) Data {

	// sudo code:
	// split each element into partitions of size 1
	// recursively merge adjacent partitions
	//   for i = leftPartIdx to rightPartIdx
	//     if leftPartHeadValue <= rightPartHeadValue
	//       copy leftPartHeadValue
	//     else: copy rightPartHeadValue; Increase InvIdx
	// copy elements back to original array

	var num = len(d) // 10

	if num == 1 {
		return d
	}

	middle := int(num / 2) // 5
	var (
		left  = make([]int, middle)     // 5
		right = make([]int, num-middle) // 5
	)
	for i := 0; i < num; i++ {
		if i < middle {
			left[i] = d[i]
		} else {
			right[i-middle] = d[i]
		}
	}

	return merge(Sort(left), Sort(right))
}

func merge(l, r []int) []int {
	result := make([]int, len(l)+len(r))

	i := 0
	for len(l) > 0 && len(r) > 0 {
		if l[0] < r[0] {
			result[i] = l[0]
			l = l[1:]
		} else {
			result[i] = r[0]
			r = r[1:]
		}
		i++
	}

	// Copy remaining items in left list if any
	for j := 0; j < len(l); j++ {
		result[i] = l[j]
		i++
	}
	// Copy remaining items in right list if any
	for j := 0; j < len(r); j++ {
		result[i] = r[j]
		i++
	}

	return result
}
