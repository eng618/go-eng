package stack_test

import (
	"fmt"
	"reflect"
	"sync"
	"testing"

	"github.com/eng618/go-eng/ds/stack"
)

func Example_basicUsage() {
	s := stack.New()

	s.Push(25)
	s.Push(1)
	s.Push(2)

	if v, err := s.Pop(); err == nil {
		fmt.Println("Pop returned", v)
	}
	// Output:
	// Pop returned 2
}

func Example_seededList() {
	s := stack.NewSeeded([]stack.Item{1, 2, 3, 4, 5})

	if v, err := s.Pop(); err == nil {
		fmt.Println("Pop returned", v)
	}
	// Output:
	// Pop returned 5
}

func TestStack_New(t *testing.T) {
	t.Parallel()
	t.Run("Create new seeded stack", func(t *testing.T) {
		t.Parallel()
		if got := stack.NewSeeded([]stack.Item{1, 2, 3, 4, 5}); !reflect.DeepEqual(got.Items, []stack.Item{1, 2, 3, 4, 5}) {
			t.Errorf("Stack.New() = %v", got)
		}
	})
	t.Run("Create new stack", func(t *testing.T) {
		t.Parallel()
		if got := stack.New(); !reflect.DeepEqual(got.Items, []stack.Item{}) {
			t.Errorf("Stack.New() = %v", got)
		}
	})
}

func TestStack_Push(t *testing.T) {
	t.Parallel()

	type args struct {
		i stack.Item
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
			t.Parallel()
			s := stack.New()
			if gotOk := s.Push(tt.args.i); gotOk != tt.wantOk {
				t.Errorf("Stack.Push() = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

//nolint:paralleltest // tests need to be sequential
func TestStack_Pop(t *testing.T) {
	s := stack.NewSeeded([]stack.Item{1, 2, 3, 4, 5})

	tests := []struct {
		name    string
		wantI   stack.Item
		wantErr bool
	}{
		{name: "Pop 5", wantI: 5, wantErr: false},
		{name: "Pop 4", wantI: 4, wantErr: false},
		{name: "Pop 3", wantI: 3, wantErr: false},
		{name: "Pop 2", wantI: 2, wantErr: false},
		{name: "Pop 1", wantI: 1, wantErr: false},
		{name: "Pop on empty stack", wantI: -1, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotI, gotErr := s.Pop()
			if gotI != tt.wantI {
				t.Errorf("Stack.Pop() gotI = %v, want %v", gotI, tt.wantI)
			}
			if (gotErr != nil) != tt.wantErr {
				t.Errorf("Stack.Pop() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}

func BenchmarkStack_Push(b *testing.B) {
	s := stack.New()
	for i := 0; i < b.N; i++ {
		s.Push(stack.Item(i))
	}
}

func BenchmarkStack_Pop(b *testing.B) {
	s := stack.NewSeeded([]stack.Item{1, 2, 3, 4, 5})
	for i := 0; i < b.N; i++ {
		s.Pop()
	}
}

func TestStack_Concurrency(t *testing.T) {
	t.Parallel()
	s := stack.New()
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			s.Push(stack.Item(i))
		}(i)
	}
	wg.Wait()

	if len(s.Items) != 1000 {
		t.Errorf("Expected stack length 1000, got %d", len(s.Items))
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			s.Pop()
		}()
	}
	wg.Wait()

	if len(s.Items) != 0 {
		t.Errorf("Expected stack length 0, got %d", len(s.Items))
	}
}
