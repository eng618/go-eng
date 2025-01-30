package sort_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/eng618/go-eng/algo/sort"
)

func ExampleStandard() {
	l := []int{1, 3, 5}
	r := []int{2, 4, 6}
	result := sort.Standard(l, r)
	fmt.Println(result)
	// Output: [1 2 3 4 5 6]
}

func TestStandard(t *testing.T) {
	t.Parallel()

	type args struct {
		l []int
		r []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "Basic1", args: args{l: []int{1}, r: []int{5}}, want: []int{1, 5}},
		{name: "Basic2", args: args{l: []int{5}, r: []int{1}}, want: []int{1, 5}},
		{name: "Basic3", args: args{l: []int{1, 2, 3, 4}, r: []int{5, 6, 7, 8, 9}}, want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{name: "Basic4", args: args{l: []int{1, 4, 7, 9}, r: []int{2, 3, 5, 6, 8}}, want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{name: "Empty right", args: args{l: []int{1, 2, 3, 4, 5}, r: []int{}}, want: []int{1, 2, 3, 4, 5}},
		{name: "Empty left", args: args{l: []int{}, r: []int{1, 2, 3, 4, 5}}, want: []int{1, 2, 3, 4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := sort.Standard(tt.args.l, tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSimple() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStandard_EdgeCases(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		l    []int
		r    []int
		want []int
	}{
		{name: "All same elements", l: []int{1, 1, 1}, r: []int{1, 1, 1}, want: []int{1, 1, 1, 1, 1, 1}},
		{name: "Negative numbers", l: []int{-3, -1, -2}, r: []int{-5, -4}, want: []int{-5, -4, -3, -2, -1}},
		{name: "Mixed positive and negative", l: []int{3, -1}, r: []int{2, -5, 4}, want: []int{-5, -1, 2, 3, 4}},
		{name: "Already sorted", l: []int{1, 2, 3}, r: []int{4, 5}, want: []int{1, 2, 3, 4, 5}},
		{name: "Reverse sorted", l: []int{5, 4}, r: []int{3, 2, 1}, want: []int{1, 2, 3, 4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := sort.Standard(tt.l, tt.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSimple() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkStandard(b *testing.B) {
	l := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	r := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}

	for i := 0; i < b.N; i++ {
		sort.Standard(l, r)
	}
}
