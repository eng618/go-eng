package search_test

import (
	"fmt"
	"testing"

	"garciaericn.com/go-eng/algo/search"
)

func ExampleBinaryLoop() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	if v, ok := search.BinaryLoop(xi, 5); ok {
		fmt.Println("Found 5 at index", v)
	}

	if v, ok := search.BinaryLoop(xi, 25); ok {
		fmt.Println("Found 25 at index", v)
	} else {
		fmt.Println("target number no found in slice")
	}
	// Output:
	// Found 5 at index 4
	// target number no found in slice
}

func ExampleBinaryRecursion() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	fmt.Println("5 is in xi =", search.BinaryRecursion(xi, 5))
	fmt.Println("25 is in xi =", search.BinaryRecursion(xi, 25))
	// Output:
	// 5 is in xi = true
	// 25 is in xi = false
}

func TestBinaryLoop(t *testing.T) {
	t.Parallel()

	type args struct {
		xi     []int
		target int
	}

	tests := []struct {
		name      string
		args      args
		wantIndex int
		wantOk    bool
	}{
		{name: "9", args: args{xi: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, target: 9}, wantIndex: 8, wantOk: true},
		{name: "1", args: args{xi: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, target: 1}, wantIndex: 0, wantOk: true},
		{name: "5", args: args{xi: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, target: 5}, wantIndex: 4, wantOk: true},
		{name: "not found", args: args{xi: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, target: 25}, wantIndex: -1, wantOk: false},
		{name: "empty slice", args: args{xi: []int{}, target: 25}, wantIndex: -1, wantOk: false},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			gotIndex, gotOk := search.BinaryLoop(tt.args.xi, tt.args.target)
			if gotIndex != tt.wantIndex {
				t.Errorf("BinarySearchForLoop() gotIndex = %v, want %v", gotIndex, tt.wantIndex)
			}
			if gotOk != tt.wantOk {
				t.Errorf("BinarySearchForLoop() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestBinaryRecursion(t *testing.T) {
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
		{name: "short1", args: args{xi: []int{1, 2, 3}, target: 1}, want: true},
		{name: "short2", args: args{xi: []int{1, 2, 3}, target: 2}, want: true},
		{name: "short3", args: args{xi: []int{1, 2, 3}, target: 3}, want: true},
		{name: "short4", args: args{xi: []int{1, 2, 3}, target: 4}, want: false},
		{
			name: "large5",
			args: args{
				xi:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
				target: 5,
			},
			want: true,
		},
		{
			name: "large50",
			args: args{
				xi:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
				target: 50,
			},
			want: false,
		},
		{
			name: "large-negative",
			args: args{
				xi:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
				target: -5,
			},
			want: false,
		},
		{name: "Should accept empty slice", args: args{xi: []int{}, target: 4}, want: false},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := search.BinaryRecursion(tt.args.xi, tt.args.target); got != tt.want {
				t.Errorf("BinarySearchRecursion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkBinaryLoop(b *testing.B) {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	target := 5
	notFoundTarget := 25

	for i := 0; i < b.N; i++ {
		if i%2 == 0 {
			search.BinaryLoop(xi, target)
		} else {
			search.BinaryLoop(xi, notFoundTarget)
		}
	}
}

func BenchmarkBinaryRecursion(b *testing.B) {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	target := 5
	notFoundTarget := 25

	for i := 0; i < b.N; i++ {
		if i%2 == 0 {
			search.BinaryRecursion(xi, target)
		} else {
			search.BinaryRecursion(xi, notFoundTarget)
		}
	}
}
