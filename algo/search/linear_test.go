package search_test

import (
	"fmt"
	"testing"

	"github.com/eng618/go-eng/algo/search"
)

func ExampleLinear() {
	exampleOne := search.Linear([]int{1, 3, 5, 7}, 5)
	exampleTwo := search.Linear([]int{1, 3, 5, 7}, 11)

	fmt.Println(exampleOne)
	fmt.Println(exampleTwo)

	// Output:
	// true
	// false
}

func TestLinear(t *testing.T) {
	t.Parallel()

	type args struct {
		xi     []int
		target int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Found", args: args{xi: []int{1, 2, 3, 4, 5}, target: 3}, want: true},
		{name: "Found first", args: args{xi: []int{1, 2, 3, 4, 5}, target: 1}, want: true},
		{name: "Found last", args: args{xi: []int{1, 2, 3, 4, 5}, target: 5}, want: true},
		{name: "Not found", args: args{xi: []int{1, 2, 3, 4, 5}, target: 25}, want: false},
		{name: "Empty slice", args: args{xi: []int{}, target: 25}, want: false},
		{name: "Negative ints", args: args{xi: []int{-25, 0, 25, 50, 100}, target: -25}, want: true},
		{name: "Negative ints", args: args{xi: []int{-100, -50, -25, 0}, target: -25}, want: true},
		{
			name: "Unsorted",
			args: args{
				xi:     []int{3, 44, 38, 5, 47, 15, 36, 26, 27, 2, 46, 4, 19, 50, 48},
				target: 46,
			},
			want: true,
		},
		{
			name: "Unsorted not found",
			args: args{
				xi:     []int{3, 44, 38, 5, 47, 15, 36, 26, 27, 2, 46, 4, 19, 50, 48},
				target: 12,
			},
			want: false,
		},
		{
			name: "Single element found",
			args: args{xi: []int{42}, target: 42},
			want: true,
		},
		{
			name: "Single element not found",
			args: args{xi: []int{42}, target: 24},
			want: false,
		},
		{
			name: "All elements same and found",
			args: args{xi: []int{7, 7, 7, 7, 7}, target: 7},
			want: true,
		},
		{
			name: "All elements same and not found",
			args: args{xi: []int{7, 7, 7, 7, 7}, target: 3},
			want: false,
		},
		{
			name: "Large slice found",
			args: args{xi: func() []int {
				xi := make([]int, 1000000)
				for i := range xi {
					xi[i] = i
				}
				return xi
			}(), target: 999999},
			want: true,
		},
		{
			name: "Large slice not found",
			args: args{xi: func() []int {
				xi := make([]int, 1000000)
				for i := range xi {
					xi[i] = i
				}
				return xi
			}(), target: 1000001},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := search.Linear(tt.args.xi, tt.args.target); got != tt.want {
				t.Errorf("Linear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkLinear(b *testing.B) {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	target := 5
	notFoundTarget := 25

	for i := 0; i < b.N; i++ {
		if i%2 == 0 {
			search.Linear(xi, target)
		} else {
			search.Linear(xi, notFoundTarget)
		}
	}
}
