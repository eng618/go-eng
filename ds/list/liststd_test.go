package list_test

import (
	"fmt"
	"testing"

	"github.com/eng618/go-eng/ds/list"
)

// Example tests.
func ExampleList() {
	l := list.NewList()
	l.PushBack(1)
	l.PushFront(0)
	l.PushBack(2)
	l.Iterate(func(value interface{}) {
		fmt.Println(value)
	})
	// Output:
	// 0
	// 1
	// 2
}

// Edge case tests.
func TestList(t *testing.T) {
	l := list.NewList()
	if l.Len() != 0 {
		t.Errorf("expected length 0, got %d", l.Len())
	}

	l.PushBack(1)
	if l.Len() != 1 {
		t.Errorf("expected length 1, got %d", l.Len())
	}

	l.PushFront(0)
	if l.Front().Value != 0 {
		t.Errorf("expected front value 0, got %v", l.Front().Value)
	}

	if l.Back().Value != 1 {
		t.Errorf("expected back value 1, got %v", l.Back().Value)
	}

	l.Remove(l.Front())
	if l.Len() != 1 {
		t.Errorf("expected length 1, got %d", l.Len())
	}

	l.Remove(l.Back())
	if l.Len() != 0 {
		t.Errorf("expected length 0, got %d", l.Len())
	}
}

// TestIterate tests the Iterate method.
func TestIterate(t *testing.T) {
	l := list.NewList()
	values := []int{1, 2, 3, 4, 5}
	for _, v := range values {
		l.PushBack(v)
	}

	var iteratedValues []int
	l.Iterate(func(value interface{}) {
		iteratedValues = append(iteratedValues, value.(int))
	})

	if len(iteratedValues) != len(values) {
		t.Errorf("expected %d values, got %d", len(values), len(iteratedValues))
	}

	for i, v := range values {
		if iteratedValues[i] != v {
			t.Errorf("expected value %d at index %d, got %d", v, i, iteratedValues[i])
		}
	}
}

// Table-driven tests.
func TestList_TableDriven(t *testing.T) {
	tests := []struct {
		name     string
		actions  func(l *list.List)
		expected []interface{}
	}{
		{
			name: "PushBack and PushFront",
			actions: func(l *list.List) {
				l.PushBack(1)
				l.PushFront(0)
				l.PushBack(2)
			},
			expected: []interface{}{0, 1, 2},
		},
		{
			name: "Remove elements",
			actions: func(l *list.List) {
				l.PushBack(1)
				l.PushBack(2)
				l.PushBack(3)
				l.Remove(l.Front())
				l.Remove(l.Back())
			},
			expected: []interface{}{2},
		},
		{
			name: "Empty list",
			actions: func(_ *list.List) {
				// No actions
			},
			expected: []interface{}{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := list.NewList()
			tt.actions(l)

			var result []interface{}
			l.Iterate(func(value interface{}) {
				result = append(result, value)
			})

			if len(result) != len(tt.expected) {
				t.Errorf("expected %d elements, got %d", len(tt.expected), len(result))
			}

			for i, v := range tt.expected {
				if result[i] != v {
					t.Errorf("expected value %v at index %d, got %v", v, i, result[i])
				}
			}
		})
	}
}

// Benchmark tests.
func BenchmarkList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := list.NewList()
		for j := 0; j < 1000; j++ {
			l.PushBack(j)
		}
		for l.Len() > 0 {
			l.Remove(l.Front())
		}
	}
}
