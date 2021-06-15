// Package stack is a simple implementation of the stack data structure.
package stack

import (
	"log"
	"sync"
)

// Item is the type a stack accepts.
type Item int

// Stack is the structure to interact with the stack package.
type Stack struct {
	items []Item
	lock  sync.RWMutex
}

// New creates an empty stack to be pushed too.
func New() *Stack {
	s := &Stack{}
	s.items = []Item{}
	log.Println("Created a new stack")
	return s
}

// NewSeeded creates a new stack with a seeded list xi.
func NewSeeded(xi []Item) *Stack {
	s := &Stack{}
	s.items = xi
	log.Println("Created a new seeded stack")
	return s
}

// Push adds the provided item to the stack.
func (s *Stack) Push(i Item) (ok bool) {
	s.lock.Lock()
	s.items = append(s.items, i)
	log.Println("Added", i, "to top of stack")
	s.lock.Unlock()
	return true
}

// Pop returns the last item placed on the stack. LIFO. Returns -1 if there is nothing on the list to return.
func (s *Stack) Pop() (i Item, ok bool) {
	if len(s.items) <= 0 {
		log.Println("There are no more items on stack")
		return -1, false
	}
	s.lock.Lock()
	i = s.items[len(s.items)-1]        // get last item
	s.items = s.items[:len(s.items)-1] // slice the last item off of slice
	s.lock.Unlock()
	return i, true
}
