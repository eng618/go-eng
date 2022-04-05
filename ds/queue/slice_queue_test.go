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

	fmt.Println(dq)

	fmt.Println(q.Length())
	q.Print()

	// Output:
	// first
	// 1
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
		queue []interface{}
	}

	tests := []struct {
		name    string
		fields  fields
		wantV   interface{}
		wantErr bool
	}{
		{name: "Empty queue", fields: fields{queue: []interface{}{}}, wantV: nil, wantErr: true},
		{name: "Last item", fields: fields{queue: []interface{}{1}}, wantV: 1, wantErr: false},
		{
			name: "Can dequeue bool",
			fields: fields{
				queue: []interface{}{true, 1, 2, "hello"},
			},
			wantV:   true,
			wantErr: false,
		},
		{
			name: "Can dequeue int",
			fields: fields{
				queue: []interface{}{2, "hello"},
			},
			wantV:   2,
			wantErr: false,
		},
		{
			name: "Can dequeue string",
			fields: fields{
				queue: []interface{}{"hello", true, 1, 2},
			},
			wantV:   "hello",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
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
	for j := 0; j < 1000000; j++ {
		q.Enqueue(j)
	}

	for i := 0; i < b.N; i++ {
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
		v interface{}
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
		tt := tt
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
	for i := 0; i < b.N; i++ {
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
}

func TestSliceQueue_Peek(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		fields    []interface{}
		wantValue interface{}
		wantErr   bool
	}{
		{name: "Empty queue", fields: []interface{}{}, wantValue: nil, wantErr: true},
		{name: "int first", fields: []interface{}{1, 2, 3, 4, 5}, wantValue: 1, wantErr: false},
		{name: "string first", fields: []interface{}{"hi", "from", "a", "test"}, wantValue: "hi", wantErr: false},
		{
			name:      "slice of int fist",
			fields:    []interface{}{[]int{1, 2, 3}, "something", true},
			wantValue: []int{1, 2, 3},
			wantErr:   false,
		},
		{name: "bool first", fields: []interface{}{true, false, 5, 4, 3, 2, 1}, wantValue: true, wantErr: false},
	}
	for _, tt := range tests {
		tt := tt
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

	for i := 0; i < b.N; i++ {
		_, _ = q.Peek()
	}
}
