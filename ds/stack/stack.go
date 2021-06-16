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

// New creates an empty stack to be pushed too.
func New() *Stack {
	s := &Stack{Items: []Item{}, lock: sync.RWMutex{}}
	s.Items = []Item{}

	return s
}

// NewSeeded creates a new stack with a seeded list xi.
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

// Pop returns the last item placed on the stack. LIFO. Returns -1 if there is nothing on the list to return.
func (s *Stack) Pop() (i Item, err error) {
	if len(s.Items) == 0 {
		return -1, errors.New("cannot pop an empty stack")
	}

	s.lock.Lock()
	i = s.Items[len(s.Items)-1]        // get last item
	s.Items = s.Items[:len(s.Items)-1] // slice the last item off of slice
	s.lock.Unlock()

	return i, nil
}
