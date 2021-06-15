package queue

import (
	"errors"

	"github.com/eng618/go-eng/ds/list"
)

// LinkedQueue is a structure used to interact with the queue package
type LinkedQueue struct {
	list list.LinkedList
}

// Dequeue is a method to get the next item in a LinkedQueue.
func (queue *LinkedQueue) Dequeue() (value interface{}, err error) {
	if queue.list.Length() == 0 {
		return nil, errors.New("queue is empty")
	}
	value, err = queue.list.PeekFront()
	if err == nil {
		if e := queue.list.Delete(0); e != nil {
			return
		}
	}
	return nil, errors.New("Dequeue exited unsuccessfully")
}

// Enqueue is a method to add an item to a LinkedQueue.
func (queue *LinkedQueue) Enqueue(value interface{}) {
	queue.list.PushBack(value)
}

// Length returns the current length of a LinkedQueue.
func (queue *LinkedQueue) Length() int {
	return queue.list.Length()
}
