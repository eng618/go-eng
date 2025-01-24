// Package anagrams is a package to identify anagrams.
package anagrams

import (
	"regexp"
	"strings"
)

// IsAnagram checks if two strings are anagrams of each other.
// This disregards punctuation, and only looks at alphanumeric values.
// An anagram is a word or phrase formed by rearranging the letters of a different word or phrase,
// typically using all the original letters exactly once.
//
// Parameters:
//   - s1: The first string to compare.
//   - s2: The second string to compare.
//
// Returns:
//   - bool: True if the strings are anagrams, false otherwise.
func IsAnagram(s1, s2 string) bool {
	cm1 := buildCharMap(s1)
	cm2 := buildCharMap(s2)

	if len(cm1) != len(cm2) {
		return false
	}

	for k, v := range cm1 {
		if _, ok := cm2[k]; !ok || v != cm2[k] {
			return false
		}
	}

	return true
}

// buildCharMap takes a string and returns a map where the keys are the
// characters in the string and the values are the counts of each character.
// It removes all non-alphanumeric characters from the string and converts
// it to lowercase before building the map.
//
// Parameters:
//
//	s - the input string to be processed.
//
// Returns:
//
//	A map[string]int where the keys are the characters from the input string
//	and the values are the counts of each character.
func buildCharMap(s string) map[string]int {
	nonAlphanumericRegex := regexp.MustCompile(`[^a-zA-Z0-9 ]+`)
	charMap := make(map[string]int)
	clean := nonAlphanumericRegex.ReplaceAllString(s, "")

	for _, v := range strings.ToLower(clean) {
		if val, ok := charMap[string(v)]; ok {
			// Increment the value of the letter in the map.
			charMap[string(v)] = val + 1
		} else {
			// Add new letter to the map, with a value of 1.
			charMap[string(v)] = 1
		}
	}
	return charMap
}
