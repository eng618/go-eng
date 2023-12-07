// Package vowels has various methods for various operations on manipulating
// vowels within a strings.
package vowels

import "strings"

// Count returns the number of vowels in the provided string.
func Count(s string) (count int) {
	for _, v := range s {
		if strings.EqualFold(string(v), "a") ||
			strings.EqualFold(string(v), "e") ||
			strings.EqualFold(string(v), "i") ||
			strings.EqualFold(string(v), "o") ||
			strings.EqualFold(string(v), "u") {
			count++
		}
	}

	return
}
