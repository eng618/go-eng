// Package anagrams is a package to identify anagrams.
package anagrams_test

import (
	"testing"

	"github.com/eng618/go-eng/algo/anagrams"
)

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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := anagrams.IsAnagram(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("IsAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}
