package stack

import (
	"fmt"
	"reflect"
	"testing"
)

func Example_basicUsage() {
	s := New()

	s.Push(25)
	s.Push(1)
	s.Push(2)
	if v, ok := s.Pop(); ok {
		fmt.Println("Pop returned", v)
	}
	// Output:
	// Pop returned 2
}

func Example_seededList() {
	s := NewSeeded([]Item{1, 2, 3, 4, 5})

	if v, ok := s.Pop(); ok {
		fmt.Println("Pop returned", v)
	}
	// Output:
	// Pop returned 5
}

func TestStack_New(t *testing.T) {
	t.Run("Create new seeded stack", func(t *testing.T) {
		if got := NewSeeded([]Item{1, 2, 3, 4, 5}); !reflect.DeepEqual(got.items, []Item{1, 2, 3, 4, 5}) {
			t.Errorf("Stack.New() = %v", got)
		}
	})
	t.Run("Create new stack", func(t *testing.T) {
		if got := New(); !reflect.DeepEqual(got.items, []Item{}) {
			t.Errorf("Stack.New() = %v", got)
		}
	})
}

func TestStack_Push(t *testing.T) {
	type args struct {
		i Item
	}
	tests := []struct {
		name   string
		args   args
		wantOk bool
	}{
		{name: "Basic push", args: args{5}, wantOk: true},
		{name: "Basic push", args: args{25}, wantOk: true},
		{name: "Basic push", args: args{1000}, wantOk: true},
		{name: "Basic push", args: args{-1}, wantOk: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			if gotOk := s.Push(tt.args.i); gotOk != tt.wantOk {
				t.Errorf("Stack.Push() = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestStack_Pop(t *testing.T) {
	s := NewSeeded([]Item{1, 2, 3, 4, 5})

	tests := []struct {
		name   string
		wantI  Item
		wantOk bool
	}{
		{name: "Pop 5", wantI: 5, wantOk: true},
		{name: "Pop 4", wantI: 4, wantOk: true},
		{name: "Pop 3", wantI: 3, wantOk: true},
		{name: "Pop 2", wantI: 2, wantOk: true},
		{name: "Pop 1", wantI: 1, wantOk: true},
		{name: "Pop on empty stack", wantI: -1, wantOk: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotI, gotOk := s.Pop()
			if gotI != tt.wantI {
				t.Errorf("Stack.Pop() gotI = %v, want %v", gotI, tt.wantI)
			}
			if gotOk != tt.wantOk {
				t.Errorf("Stack.Pop() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
