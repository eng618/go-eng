// Package binary is a simple implementation of binary search algorithm
package binary

import (
	"log"
)

// BinarySearchForLoop is a binary search implementation to search an slice of type int
// BinarySearchForLoop is an algorithm that uses the binary search method to search a given slice.
// For this example we will assume the slice is sorted.
//
// https://goplay.space/#TQBxLBIrt-w
//
// https://play.golang.org/p/TQBxLBIrt-w
func BinarySearchForLoop(xi []int, target int) (index int, ok bool) {
	min, max := 0, len(xi)-1
	// guess is an int so it will always be the truncated intiger number.
	var guess int

	numGuess := 0

	for min <= max {
		numGuess++
		guess = (min + max) / 2

		if xi[guess] == target {
			log.Println("found target number in", numGuess, "tries")
			return guess, true
		} else if xi[guess] < target {
			min = guess + 1
		} else {
			max = guess - 1
		}
	}
	log.Println("made", numGuess, "guesses before I realise the number was not present")
	return -1, false
}
