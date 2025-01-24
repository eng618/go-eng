package queue

import (
	"errors"
	"fmt"
)

// SliceQueue represents a queue implemented using a slice.
type SliceQueue struct {
	// queue holds the elements of the queue.
	queue []interface{}
	// length is the current number of elements in the queue.
	length int
}

// Dequeue removes and returns the front element of the queue.
// If the queue is empty, it returns an error.
// Returns:
// - v: the front element of the queue.
// - err: an error if the queue is empty.
func (q *SliceQueue) Dequeue() (v interface{}, err error) {
	if q.length == 0 {
		return nil, errors.New("attempted to dequeue on an empty queue")
	}

	v = q.queue[0]
	q.queue = q.queue[1:]
	q.length--

	return v, nil
}

// Enqueue is a method to add an item to a SliceQueue.
func (q *SliceQueue) Enqueue(v interface{}) (ok bool) {
	q.queue = append(q.queue, v)
	q.length++

	return true
}

// Peek shows the head or next item in the queue.
func (q *SliceQueue) Peek() (value interface{}, err error) {
	if q.length == 0 {
		return nil, errors.New("there are no items in this queue")
	}

	return q.queue[0], nil
}

// Length returns the number of elements currently in the queue.
func (q *SliceQueue) Length() int {
	return q.length
}

// Print prints all the elements in the queue to the standard output.
// If the queue is empty, it prints "Queue is empty".
func (q *SliceQueue) Print() {
	if q.length == 0 {
		fmt.Println("Queue is empty")
	}

	for _, v := range q.queue {
		fmt.Println(v)
	}
}
