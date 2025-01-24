// Package anagrams is a package to identify anagrams.
package anagrams_test

import (
	"fmt"
	"testing"

	"github.com/eng618/go-eng/algo/anagrams"
)

// ExampleIsAnagram provides an example usage of the IsAnagram function.
func ExampleIsAnagram() {
	result := anagrams.IsAnagram("listen", "silent")
	fmt.Println(result)
	// Output: true
}

func TestIsAnagram(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Positive case", args: args{s1: "rail safety", s2: "fairy tales"}, want: true},
		{name: "Negative case", args: args{s1: "Hi there", s2: "Bye there"}, want: false},
		{name: "Positive with various punctuation", args: args{s1: "RAIL! SAFETY!", s2: "fairy tales"}, want: true},
		{name: "hello", args: args{s1: "hello", s2: "llohe"}, want: true},
		{name: "Mismatched letters", args: args{s1: "hello", s2: "heloo"}, want: false},
		{name: "listen_silent", args: args{s1: "listen", s2: "silent"}, want: true},
		{name: "triangle_integral", args: args{s1: "triangle", s2: "integral"}, want: true},
		{name: "apple_pale", args: args{s1: "apple", s2: "pale"}, want: false},
		{name: "empty strings", args: args{s1: "", s2: ""}, want: true},
		{name: "single character", args: args{s1: "a", s2: "a"}, want: true},
		{name: "two characters", args: args{s1: "ab", s2: "ba"}, want: true},
		{name: "mismatched two characters", args: args{s1: "ab", s2: "bc"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := anagrams.IsAnagram(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("IsAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkIsAnagram benchmarks the IsAnagram function.
func BenchmarkIsAnagram(b *testing.B) {
	for i := 0; i < b.N; i++ {
		anagrams.IsAnagram("listen", "silent")
	}
}
