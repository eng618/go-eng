package list

import (
	"container/list"
)

// List is a wrapper around the container/list.List.
type List struct {
	list *list.List
}

// New creates a new List.
func NewList() *List {
	return &List{list: list.New()}
}

// PushBack adds an element to the end of the list.
func (l *List) PushBack(value interface{}) {
	l.list.PushBack(value)
}

// PushFront adds an element to the front of the list.
func (l *List) PushFront(value interface{}) {
	l.list.PushFront(value)
}

// Front returns the first element of the list.
func (l *List) Front() *list.Element {
	return l.list.Front()
}

// Back returns the last element of the list.
func (l *List) Back() *list.Element {
	return l.list.Back()
}

// Remove removes an element from the list.
func (l *List) Remove(e *list.Element) interface{} {
	return l.list.Remove(e)
}

// Len returns the number of elements in the list.
func (l *List) Len() int {
	return l.list.Len()
}

// Iterate iterates through the list and applies the given function to each element.
func (l *List) Iterate(f func(value interface{})) {
	for e := l.list.Front(); e != nil; e = e.Next() {
		f(e.Value)
	}
}
