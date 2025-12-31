package sorting_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/eng618/go-eng/algo/sorting"
)

func ExampleBubbleSort() {
	input := []int{5, 3, 8, 4, 2}
	sorted := sorting.BubbleSort(input)
	fmt.Println(sorted)
	// Output: [2 3 4 5 8]
}

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"sorted", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"reverse", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"unsorted", []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}, []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}},
		{"empty", []int{}, []int{}},
		{"single element", []int{1}, []int{1}},
		{"all same", []int{2, 2, 2, 2}, []int{2, 2, 2, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sorting.BubbleSort(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("BubbleSort(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	input := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	for i := 0; i < b.N; i++ {
		sorting.BubbleSort(input)
	}
}
