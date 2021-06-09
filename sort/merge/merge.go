// Package merge uses the merge sort algorithm.
// The runtime of merge sort is at best, at worst, and at average always O(n * logn)
//
// See https://visualgo.net/en/sorting for a visual example of merge sort.
package merge

type Data []int

// Sort takes the provided data (slice of int) and applies the merge sort algorithm, to sort the data.
func Sort(d Data) Data {
	var num = len(d)

	if num <= 1 {
		return d
	}

	middle := int(num / 2)
	var (
		left  = make([]int, middle)
		right = make([]int, num-middle)
	)
	// split data into 2 halves
	for i := 0; i < num; i++ {
		if i < middle {
			left[i] = d[i]
		} else {
			right[i-middle] = d[i]
		}
	}

	// recursively merge sorted sides
	return merge(Sort(left), Sort(right))
}

func merge(l, r Data) Data {
	result := make([]int, len(l)+len(r))

	// Add items to result until either side is empty.
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
