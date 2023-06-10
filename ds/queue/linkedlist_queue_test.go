package queue_test

import (
	"reflect"
	"testing"

	"github.com/eng618/go-eng/ds/queue"
)

//nolint:dupl // This code is duplicated to test slice & linked list queues
func TestLinkedQueue_Dequeue(t *testing.T) {
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
			q := &queue.LinkedQueue{}
			for _, v := range tt.fields {
				q.Enqueue(v)
			}
			gotValue, err := q.Dequeue()
			if (err != nil) != tt.wantErr {
				t.Errorf("LinkedQueue.Dequeue() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if !reflect.DeepEqual(gotValue, tt.wantValue) {
				t.Errorf("LinkedQueue.Dequeue() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func BenchmarkLinkedQueue_Dequeue(b *testing.B) {
	q := &queue.LinkedQueue{}
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

func TestLinkedQueue_Enqueue(t *testing.T) {
	t.Parallel()

	type args struct {
		value interface{}
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "int", args: args{value: 1}},
		{name: "string", args: args{value: "Hi!"}},
		{name: "[]int", args: args{value: []int{1, 2, 3, 4, 5}}},
		{name: "map", args: args{value: map[int]string{1: "hello", 2: "gophers"}}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			q := &queue.LinkedQueue{}
			q.Enqueue(tt.args.value)
		})
	}
}

func TestLinkedQueue_Length(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		fields []interface{}
		want   int
	}{
		{name: "empty", fields: []interface{}{}, want: 0},
		{name: "one", fields: []interface{}{1}, want: 1},
		{name: "two", fields: []interface{}{1, 2}, want: 2},
		{name: "three", fields: []interface{}{1, 2, 3}, want: 3},
		{
			name:   "twenty-five",
			fields: []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25},
			want:   25,
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			q := &queue.LinkedQueue{}

			for _, v := range tt.fields {
				q.Enqueue(v)
			}
			if got := q.Length(); got != tt.want {
				t.Errorf("LinkedQueue.Length() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedQueue_Peek(t *testing.T) {
	t.Parallel()

	q := &queue.LinkedQueue{}
	q.Enqueue("head")
	q.Enqueue(1)
	q.Enqueue("tail")

	gotValue, err := q.Peek()
	if err != nil {
		t.Errorf("LinkedQueue.Peek() error = %v, wantErr %v", err, false)

		return
	}

	if !reflect.DeepEqual(gotValue, "head") {
		t.Errorf("LinkedQueue.Peek() = %v, want %v", gotValue, "head")
	}

	q2 := &queue.LinkedQueue{}
	if _, err := q2.Peek(); err == nil {
		t.Errorf("LinkedQueue.Peek() error = %v, wantErr %v", err, true)
	}
}
