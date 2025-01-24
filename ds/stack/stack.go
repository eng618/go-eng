// Package stack is a simple implementation of the stack data structure.
package stack

import (
	"errors"
	"sync"
)

// Item is the type a stack accepts.
type Item int

// Stack is the structure to interact with the stack package.
type Stack struct {
	Items []Item
	lock  sync.RWMutex
}

// New creates and returns a new instance of Stack with initialized fields.
// It initializes the Items slice and the lock for concurrent access control.
func New() *Stack {
	s := &Stack{Items: []Item{}, lock: sync.RWMutex{}}
	s.Items = []Item{}

	return s
}

// NewSeeded creates a new Stack instance pre-seeded with the provided items.
// It initializes the stack with the given slice of items and a read-write mutex for concurrency safety.
//
// Parameters:
//
//	xi []Item - A slice of items to initialize the stack with.
//
// Returns:
//
//	*Stack - A pointer to the newly created Stack instance.
//
// - err: An error if the stack is empty, otherwise nil.
func NewSeeded(xi []Item) *Stack {
	s := &Stack{Items: xi, lock: sync.RWMutex{}}

	return s
}

// Push adds the provided item to the stack.
func (s *Stack) Push(i Item) (ok bool) {
	s.lock.Lock()
	s.Items = append(s.Items, i)
	s.lock.Unlock()

	return true
}

// Pop removes and returns the top item from the stack.
// If the stack is empty, it returns an error.
// It locks the stack during the operation to ensure thread safety.
//
// Returns:
// - Item: The item that was removed from the stack.
// - error: An error if the stack is empty.
func (s *Stack) Pop() (i Item, err error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if len(s.Items) == 0 {
		return -1, errors.New("cannot pop an empty stack")
	}

	i = s.Items[len(s.Items)-1]        // get last item
	s.Items = s.Items[:len(s.Items)-1] // slice the last item off of slice

	return i, nil
}
