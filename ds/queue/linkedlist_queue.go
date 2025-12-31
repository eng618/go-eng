package queue

import (
	"errors"

	"github.com/eng618/go-eng/ds/list"
)

// LinkedQueue is a structure used to interact with the queue package.
type LinkedQueue struct {
	list list.LinkedList
}

// Dequeue is a method to get the next item in a LinkedQueue.
func (q *LinkedQueue) Dequeue() (value interface{}, err error) {
	if q.list.Length() == 0 {
		return nil, errors.New("queue is empty")
	}

	value, err = q.list.PeekFront()

	if err == nil {
		if e := q.list.Delete(0); e == nil {
			return value, err
		}
	}

	// This should never happen
	return nil, errors.New("Dequeue exited unsuccessfully")
}

// Enqueue is a method to add an item to a LinkedQueue.
func (q *LinkedQueue) Enqueue(value interface{}) {
	q.list.PushBack(value)
}

// Peek shows the head or next item in the queue.
func (q *LinkedQueue) Peek() (value interface{}, err error) {
	if q.list.Length() == 0 {
		return nil, errors.New("queue is empty")
	}

	return q.list.Head(), nil
}

// Length returns the current length of a LinkedQueue.
func (q *LinkedQueue) Length() int {
	return q.list.Length()
}
