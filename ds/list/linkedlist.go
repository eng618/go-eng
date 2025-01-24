// Package list is a simple implementation of the linked list data structure.
package list

import (
	"errors"
	"fmt"
	"sync"
)

// Node is the structure that makes up a single Node within a linked list.
type Node struct {
	Data interface{}
	Next *Node
}

// LinkedList is the structure of a linked list.
type LinkedList struct {
	size int
	head *Node
	tail *Node
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
		l.head = l.head.Next
		l.size--

		if l.head == nil || l.head.Next == nil {
			l.tail = nil
		}

		return nil
	default:
		current := l.head
		for i := 0; i < position-1; i++ {
			current = current.Next
		}

		current.Next = current.Next.Next
		l.size--
	}

	return nil
}

// Display prints the elements of the linked list in a readable format.
// Each element is followed by an arrow ("->"), and the last element is followed by a newline.
func (l *LinkedList) Display() {
	n := l.head

	for n != nil {
		if n.Next == nil {
			fmt.Printf("%v ->\n", n.Data)
		} else {
			fmt.Printf("%v -> ", n.Data)
		}

		n = n.Next
	}
}

// New creates and returns a new instance of a LinkedList with initialized values.
// The returned LinkedList has a size of 0, and both head and tail pointers are set to nil.
// A mutex is also initialized for concurrent access control.
func New() LinkedList {
	return LinkedList{size: 0, head: nil, tail: nil, mu: sync.Mutex{}}
}

// NewSeeded creates a new LinkedList with an initial seed value.
// The seed value is stored in the first node of the list.
// It returns a LinkedList with size 1, where both head and tail point to the seed node.
// The LinkedList is initialized with a mutex for concurrent access.
//
// Parameters:
// - seed: The initial value to be stored in the LinkedList.
//
// Returns:
// - LinkedList: A new LinkedList containing the seed value.
func NewSeeded(seed interface{}) LinkedList {
	n := &Node{Data: seed, Next: nil}
	return LinkedList{size: 1, head: n, tail: n, mu: sync.Mutex{}}
}

// PeekBack returns the data of the last element in the linked list.
// If the linked list is empty, it returns an error indicating that the list is empty.
func (l *LinkedList) PeekBack() (interface{}, error) {
	if l.head == nil {
		return nil, errors.New("cannot find Back value in an empty linked list")
	}

	return l.tail.Data, nil
}

// PeekFront returns the data of the front node in the linked list without removing it.
// If the linked list is empty, it returns an error indicating that the front value cannot be found.
func (l *LinkedList) PeekFront() (interface{}, error) {
	if l.head == nil {
		return nil, errors.New("cannot find Front value in an empty linked list")
	}

	return l.head.Data, nil
}

// PushBack adds the supplied node to the end of a LinkedList.
func (l *LinkedList) PushBack(v interface{}) {
	n := &Node{Data: v, Next: nil}

	if l.head == nil {
		l.head = n
		l.tail = n
		l.size++
	} else {
		l.tail.Next = n
		l.tail = n
		l.size++
	}
}

// PushFront adds the supplied node to the beginning of a LinkedList.
func (l *LinkedList) PushFront(data interface{}) {
	n := &Node{Data: data, Next: nil}

	if l.head == nil {
		l.head = n
		l.tail = n
		l.size++
	} else {
		n.Next = l.head
		l.head = n
		l.size++
	}
}

// Remove will removed the first occurrence of the provided data from the list, if present.
func (l *LinkedList) Remove(data interface{}) error {
	if l.size == 0 {
		return errors.New("cannot delete node from empty list")
	}

	if l.head.Data == data {
		l.head = l.head.Next
		l.size--

		if l.head == nil || l.head.Next == nil {
			l.tail = nil
		}

		return nil
	}

	n := l.head
	for i := 0; i < l.Length()-1; i++ {
		if n.Next.Data == data {
			n.Next = n.Next.Next
			l.size--

			return nil
		}

		n = n.Next
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

	var prev *Node

	for n != nil {
		temp := n.Next
		n.Next = prev
		prev = n
		n = temp
	}

	l.head = prev
}

// If the list is empty, it returns nil.
// Tail returns the data of the last node (tail) in the linked list.
// If the list is empty, it returns nil.
func (l *LinkedList) Tail() interface{} {
	return l.tail.Data
}

// Head returns the value of the current head.
// Head returns the data stored in the head node of the linked list.
// If the list is empty, it returns nil.
func (l *LinkedList) Head() interface{} {
	return l.head.Data
}

// Length returns the number of elements in the linked list.
func (l *LinkedList) Length() int {
	return l.size
}
