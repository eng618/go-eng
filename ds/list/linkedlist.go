// Package list is a simple implementation of the linked list data structure.
package list

import (
	"errors"
	"fmt"
)

// node is the structure that makes up a single node within a linked list.
type node struct {
	data interface{}
	next *node
}

// LinkedList is the structure of a linked list.
type LinkedList struct {
	Length int
	head   *node
	tail   *node
}

func New() LinkedList {
	return LinkedList{Length: 0, head: nil, tail: nil}
}

// Len returns the Length of the provided LinkedList.
func (l *LinkedList) Len() int {
	return l.Length
}

// PushBack adds the supplied node to the end of a LinkedList.
func (l *LinkedList) PushBack(v interface{}) {
	n := &node{data: v}

	if l.head == nil {
		l.head = n
		l.tail = n
		l.Length++
	} else {
		l.tail.next = n
		l.tail = n
		l.Length++
	}
}

// PushFront adds the supplied node to the beginning of a LinkedList.
func (l *LinkedList) PushFront(data interface{}) {
	n := &node{data: data}

	if l.head == nil {
		l.head = n
		l.tail = n
		l.Length++
	} else {
		n.next = l.head
		l.head = n
		l.Length++
	}
}

// Delete removes the node with the provided key from a linkedlist.
func (l *LinkedList) Delete(position int) (ok bool, err error) {

	if l.Length == 0 {
		return false, errors.New("Cannot delete node from empty list")
	}
	if position >= l.Length {
		return false, errors.New("Possition outside of range.")
	}

	switch position {
	case 0:
		l.head = l.head.next
		l.Length--
		if l.head == nil {
			l.tail = nil
		}
		return true, nil
	// case l.Length - 1:
	// 	// Tail do someting different
	// 	fmt.Println("this is interesting")
	default:
		current := l.head
		for i := 0; i < position-1; i++ {
			current = current.next
		}
		current.next = current.next.next
		l.Length--
	}

	return true, nil
}

// Display is a helper to print a visual representation of the linked list to the console.
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

// Front returns the front (or first) node of a linked list.
func (l *LinkedList) Front() (interface{}, error) {
	if l.head == nil {
		return 0, fmt.Errorf("Cannot find Front value in an empty linked list")
	}
	return l.head.data, nil
}

// Back returns the back (or last) node of a linked list.
func (l *LinkedList) Back() (interface{}, error) {
	if l.head == nil {
		return 0, fmt.Errorf("Cannot find Back value in an empty linked list")
	}
	return l.tail.data, nil
}

// Reverse takes the linked list and reverses all off the values within it.
func (l *LinkedList) Reverse() {
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
