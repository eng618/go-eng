package binary

import (
	"fmt"
	"testing"
)

func ExampleBinarySearchForLoop() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	if v, ok := BinarySearchForLoop(xi, 5); ok {
		fmt.Println("Found 5 at index", v)
	}

	if v, ok := BinarySearchForLoop(xi, 25); ok {
		fmt.Println("Found 25 at index", v)
	} else {
		fmt.Println("target number no found in slice")
	}
	// Output:
	// Found 5 at index 4
	// target number no found in slice
}

func TestBinarySearchForLoop(t *testing.T) {
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIndex, gotOk := BinarySearchForLoop(tt.args.xi, tt.args.target)
			if gotIndex != tt.wantIndex {
				t.Errorf("BinarySearchForLoop() gotIndex = %v, want %v", gotIndex, tt.wantIndex)
			}
			if gotOk != tt.wantOk {
				t.Errorf("BinarySearchForLoop() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
