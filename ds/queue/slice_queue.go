package queue

import (
	"errors"
	"fmt"
)

// SliceQueue is the structure to used to create a new queue.
// Once you have a SliceQueue, you can begin to use all the methods associated
// with it.
type SliceQueue struct {
	queue  []interface{}
	length int
}

// Dequeue is a method to get the next item in a SliceQueue.
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

// Length returns the current length of a SliceQueue.
func (q *SliceQueue) Length() int {
	return q.length
}

// Print allows you to print all the items in a SliceQueue to the screen.
func (q *SliceQueue) Print() {
	if q.length == 0 {
		fmt.Println("Queue is empty")
	}

	for _, v := range q.queue {
		fmt.Println(v)
	}
}
