// Package list is a simple implementation of the linked list data structure.
package list

import (
	"errors"
	"fmt"
	"sync"
)

// node is the structure that makes up a single node within a linked list.
type node struct {
	data interface{}
	next *node
}

// LinkedList is the structure of a linked list.
type LinkedList struct {
	size int
	head *node
	tail *node
	mu   sync.Mutex
}

// Delete removes the node at the specified position in the linked list.
// It returns an error if the list is empty or if the position is out of range.
//
// Parameters:
//
//	position (int): The zero-based index of the node to delete.
//
// Returns:
//
//	error: An error if the list is empty or if the position is out of range, otherwise nil.
func (l *LinkedList) Delete(position int) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.size == 0 {
		return errors.New("cannot delete node from empty list")
	}

	if position >= l.size {
		return errors.New("position outside of range")
	}

	switch position {
	case 0:
		l.head = l.head.next
		l.size--

		if l.head == nil || l.head.next == nil {
			l.tail = nil
		}

		return nil
	default:
		current := l.head
		for i := 0; i < position-1; i++ {
			current = current.next
		}

		current.next = current.next.next
		l.size--
	}

	return nil
}

// Display prints the elements of the linked list in a readable format.
// Each element is followed by an arrow ("->"), and the last element is followed by a newline.
func (l *LinkedList) Display() {
	n := l.head

	for n != nil {
		if n.next == nil {
			fmt.Printf("%v ->\n", n.data)
		} else {
			fmt.Printf("%v -> ", n.data)
		}

		n = n.next
	}
}

// New creates and returns a new instance of a LinkedList with initialized values.
// The returned LinkedList has a size of 0, and both head and tail pointers are set to nil.
// A mutex is also initialized for concurrent access control.
func New() LinkedList {
	return LinkedList{size: 0, head: nil, tail: nil, mu: sync.Mutex{}}
}

// NewSeeded creates a new LinkedList with a single element (seed).
// The seed parameter is the initial data for the first node in the list.
// It returns a LinkedList with size 1, the head node containing the seed data,
// and an initialized mutex for thread-safe operations.
func NewSeeded(seed interface{}) LinkedList {
	return LinkedList{size: 1, head: &node{data: seed, next: nil}, tail: nil, mu: sync.Mutex{}}
}

// PeekBack returns the data of the last element in the linked list.
// If the linked list is empty, it returns an error indicating that the list is empty.
func (l *LinkedList) PeekBack() (interface{}, error) {
	if l.head == nil {
		return nil, errors.New("cannot find Back value in an empty linked list")
	}

	return l.tail.data, nil
}

// PeekFront returns the data of the front node in the linked list without removing it.
// If the linked list is empty, it returns an error indicating that the front value cannot be found.
func (l *LinkedList) PeekFront() (interface{}, error) {
	if l.head == nil {
		return nil, errors.New("cannot find Front value in an empty linked list")
	}

	return l.head.data, nil
}

// PushBack adds the supplied node to the end of a LinkedList.
func (l *LinkedList) PushBack(v interface{}) {
	n := &node{data: v, next: nil}

	if l.head == nil {
		l.head = n
		l.tail = n
		l.size++
	} else {
		l.tail.next = n
		l.tail = n
		l.size++
	}
}

// PushFront adds the supplied node to the beginning of a LinkedList.
func (l *LinkedList) PushFront(data interface{}) {
	n := &node{data: data, next: nil}

	if l.head == nil {
		l.head = n
		l.tail = n
		l.size++
	} else {
		n.next = l.head
		l.head = n
		l.size++
	}
}

// Remove will removed the first occurrence of the provided data from the list, if present.
func (l *LinkedList) Remove(data interface{}) error {
	if l.size == 0 {
		return errors.New("cannot delete node from empty list")
	}

	if l.head.data == data {
		l.head = l.head.next
		l.size--

		if l.head == nil || l.head.next == nil {
			l.tail = nil
		}

		return nil
	}

	n := l.head
	for i := 0; i < l.Length()-1; i++ {
		if n.next.data == data {
			n.next = n.next.next
			l.size--

			return nil
		}

		n = n.next
	}

	return fmt.Errorf("unable to find %v in list", data)
}

// Reverse reverses the linked list in place. It locks the list to ensure
// thread safety, then iterates through the nodes, reversing the direction
// of the next pointers. After the loop, the head and tail pointers are
// updated to reflect the new order of the list.
func (l *LinkedList) Reverse() {
	l.mu.Lock()
	defer l.mu.Unlock()

	n := l.head
	l.tail = l.head

	var prev *node

	for n != nil {
		temp := n.next
		n.next = prev
		prev = n
		n = temp
	}

	l.head = prev
}

// If the list is empty, it returns nil.
// Tail returns the data of the last node (tail) in the linked list.
// If the list is empty, it returns nil.
func (l *LinkedList) Tail() interface{} {
	return l.tail.data
}

// Head returns the value of the current head.
// Head returns the data stored in the head node of the linked list.
// If the list is empty, it returns nil.
func (l *LinkedList) Head() interface{} {
	return l.head.data
}

// Length returns the number of elements in the linked list.
func (l *LinkedList) Length() int {
	return l.size
}
