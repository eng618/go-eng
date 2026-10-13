package queue_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/eng618/go-eng/ds/queue"
)

func Example() {
	q := queue.SliceQueue{}

	q.Enqueue("first")
	q.Enqueue("second")

	dq, err := q.Dequeue()
	if err != nil {
		fmt.Println("Dequeue err:", err)
	}

	fmt.Println("Just dequeued:", dq)

	fmt.Println("Current length is:", q.Length())
	q.Print()

	// Output:
	// Just dequeued: first
	// Current length is: 1
	// second
}

func ExampleSliceQueue_Dequeue() {
	q := queue.SliceQueue{}

	q.Enqueue("first")
	q.Enqueue("second")
	q.Enqueue("third")

	first, err := q.Dequeue()
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(first)

	second, err := q.Dequeue()
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(second)

	third, err := q.Dequeue()
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(third)

	fourth, err := q.Dequeue()
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(fourth)

	q.Print()

	// Output:
	// first
	// second
	// third
	// Error: attempted to dequeue on an empty queue
	// <nil>
	// Queue is empty
}

func TestSliceQueue_Dequeue(t *testing.T) {
	t.Parallel()

	type fields struct {
		queue []any
	}

	tests := []struct {
		name    string
		fields  fields
		wantV   any
		wantErr bool
	}{
		{name: "Empty queue", fields: fields{queue: []any{}}, wantV: nil, wantErr: true},
		{name: "Last item", fields: fields{queue: []any{1}}, wantV: 1, wantErr: false},
		{
			name: "Can dequeue bool",
			fields: fields{
				queue: []any{true, 1, 2, "hello"},
			},
			wantV:   true,
			wantErr: false,
		},
		{
			name: "Can dequeue int",
			fields: fields{
				queue: []any{2, "hello"},
			},
			wantV:   2,
			wantErr: false,
		},
		{
			name: "Can dequeue string",
			fields: fields{
				queue: []any{"hello", true, 1, 2},
			},
			wantV:   "hello",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			q := queue.SliceQueue{}
			for _, val := range tt.fields.queue {
				q.Enqueue(val)
			}
			gotV, err := q.Dequeue()
			if (err != nil) != tt.wantErr {
				t.Errorf("SliceQueue.Dequeue() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if !reflect.DeepEqual(gotV, tt.wantV) {
				t.Errorf("SliceQueue.Dequeue() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func BenchmarkSliceQueue_Dequeue(b *testing.B) {
	q := queue.SliceQueue{}
	for j := range 1000000 {
		q.Enqueue(j)
	}

	for i := range b.N {
		if i%2 == 0 {
			q.Enqueue(i)
			_, _ = q.Dequeue()
		} else {
			_, _ = q.Dequeue()
		}
	}
}

func ExampleSliceQueue_Enqueue() {
	q := queue.SliceQueue{}

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue("hello")
	q.Enqueue(true)

	q.Print()

	// Output:
	// 1
	// 2
	// hello
	// true
}

func TestSliceQueue_Enqueue(t *testing.T) {
	t.Parallel()

	type args struct {
		v any
	}

	tests := []struct {
		name   string
		q      queue.SliceQueue
		args   args
		wantOk bool
	}{
		{name: "Should accept int", q: queue.SliceQueue{}, args: args{1}, wantOk: true},
		{name: "Should accept bool", q: queue.SliceQueue{}, args: args{true}, wantOk: true},
		{name: "Should accept string", q: queue.SliceQueue{}, args: args{"hello queue"}, wantOk: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if gotOk := tt.q.Enqueue(tt.args.v); gotOk != tt.wantOk {
				t.Errorf("SliceQueue.Enqueue() = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func BenchmarkSliceQueue_Enqueue(b *testing.B) {
	q := queue.SliceQueue{}
	for range b.N {
		q.Enqueue(1)
	}
}

func ExampleSliceQueue_Peek() {
	q := &queue.SliceQueue{}

	q.Enqueue("boo")
	q.Enqueue("who?")

	if val, err := q.Peek(); err != nil {
		fmt.Println("That wasn't scary...err:", err)
	} else {
		fmt.Println(val, "AHHHH!!!!!!")
	}

	// Output:
	// boo AHHHH!!!!!!
}

func TestSliceQueue_Peek(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		fields    []any
		wantValue any
		wantErr   bool
	}{
		{name: "Empty queue", fields: []any{}, wantValue: nil, wantErr: true},
		{name: "int first", fields: []any{1, 2, 3, 4, 5}, wantValue: 1, wantErr: false},
		{name: "string first", fields: []any{"hi", "from", "a", "test"}, wantValue: "hi", wantErr: false},
		{
			name:      "slice of int fist",
			fields:    []any{[]int{1, 2, 3}, "something", true},
			wantValue: []int{1, 2, 3},
			wantErr:   false,
		},
		{name: "bool first", fields: []any{true, false, 5, 4, 3, 2, 1}, wantValue: true, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			q := &queue.SliceQueue{}
			for _, v := range tt.fields {
				q.Enqueue(v)
			}
			gotValue, err := q.Peek()
			if (err != nil) != tt.wantErr {
				t.Errorf("SliceQueue.Peek() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if !reflect.DeepEqual(gotValue, tt.wantValue) {
				t.Errorf("SliceQueue.Peek() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func BenchmarkSliceQueue_Peek(b *testing.B) {
	q := queue.SliceQueue{}
	q.Enqueue("peek benchmark")

	for range b.N {
		_, _ = q.Peek()
	}
}
