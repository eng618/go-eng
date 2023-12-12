// Package anagrams is a package to identify anagrams.
package anagrams

import (
	"fmt"
	"regexp"
	"strings"
)

// IsAnagram determines if two strings are anagrams of each other, and returns a bool.
// This disregards punctuation, and only looks at alphanumeric values.
func IsAnagram(s1, s2 string) bool {
	// strings.sort
	cm1 := buildCharMap(s1)
	cm2 := buildCharMap(s2)

	if len(cm1) != len(cm2) {
		return false
	}

	for k, v := range cm1 {
		if v != cm2[k] {
			return false
		}
	}

	return true
}

// buildCharMap helper function to create a character map of the given string.
func buildCharMap(s string) map[string]int {
	nonAlphanumericRegex := regexp.MustCompile(`[^a-zA-Z0-9 ]+`)
	charMap := make(map[string]int)
	clean := nonAlphanumericRegex.ReplaceAllString(s, "")

	for _, v := range strings.ToLower(clean) {
		if val, ok := charMap[string(v)]; ok {
			fmt.Println("ENG: adding to letter:", string(v))
			charMap[string(v)] = val + 1
		} else {
			fmt.Println("ENG: adding new letter:", string(v))
			charMap[string(v)] = 1
		}
	}
	return charMap
}
