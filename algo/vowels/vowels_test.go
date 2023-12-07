// Package vowels has various methods for various operations on manipulating
// vowels within a strings.
package vowels_test

import (
	"testing"

	"github.com/eng618/go-eng/algo/vowels"
)

func TestCount(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name      string
		args      args
		wantCount int
	}{
		{name: "All vowels", args: args{s: "aeiou"}, wantCount: 5},
		{name: "All vowels capitalized", args: args{s: "AEIOU"}, wantCount: 5},
		{name: "Hello", args: args{s: "Hello"}, wantCount: 2},
		{name: "Random with matching", args: args{s: "abcdefghijklmnopqrstuvwxyz"}, wantCount: 5},
		{name: "Random with no matching", args: args{s: "bcdfghjkl"}, wantCount: 0},
		{name: "String 1", args: args{s: "Hi There!"}, wantCount: 3},
		{name: "String 2", args: args{s: "Why do you ask?"}, wantCount: 4},
		{name: "String 3", args: args{s: "Why?"}, wantCount: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCount := vowels.Count(tt.args.s); gotCount != tt.wantCount {
				t.Errorf("Count() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}
