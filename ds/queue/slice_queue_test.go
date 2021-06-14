package queue

import (
	"fmt"
	"reflect"
	"testing"
)

func Example() {
	q := SliceQueue{}

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

func ExampleSliceQueue_Enqueue() {
	q := SliceQueue{}

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
	type args struct {
		v interface{}
	}
	tests := []struct {
		name   string
		q      SliceQueue
		args   args
		wantOk bool
	}{
		{name: "Should accept int", q: SliceQueue{}, args: args{1}, wantOk: true},
		{name: "Should accept bool", q: SliceQueue{}, args: args{true}, wantOk: true},
		{name: "Should accept string", q: SliceQueue{}, args: args{"hello queue"}, wantOk: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOk := tt.q.Enqueue(tt.args.v); gotOk != tt.wantOk {
				t.Errorf("SliceQueue.Enqueue() = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func BenchmarkSliceQueue_Enqueue(b *testing.B) {
	q := SliceQueue{}
	for i := 0; i < b.N; i++ {
		q.Enqueue(1)
	}
}

func ExampleSliceQueue_Dequeue() {
	q := SliceQueue{}

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
	// Error: Attempted to dequeue on an empty queue
	// <nil>
	// Queue is empty
}

func TestSliceQueue_Dequeue(t *testing.T) {
	type fields struct {
		queue  []interface{}
		length int
	}
	tests := []struct {
		name    string
		fields  fields
		wantV   interface{}
		wantErr bool
	}{
		{name: "Empty queue", fields: fields{queue: []interface{}{}, length: 0}, wantV: nil, wantErr: true},
		{name: "Last item", fields: fields{queue: []interface{}{1}, length: 1}, wantV: 1, wantErr: false},
		{
			name: "Can dequeue bool",
			fields: fields{
				queue:  []interface{}{true, 1, 2, "hello"},
				length: 4},
			wantV:   true,
			wantErr: false,
		},
		{
			name: "Can dequeue int",
			fields: fields{
				queue:  []interface{}{2, "hello"},
				length: 2},
			wantV:   2,
			wantErr: false,
		},
		{
			name: "Can dequeue string",
			fields: fields{
				queue:  []interface{}{"hello", true, 1, 2},
				length: 4},
			wantV:   "hello",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &SliceQueue{
				queue:  tt.fields.queue,
				length: tt.fields.length,
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
	q := SliceQueue{}
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
