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
	size int
	head *node
	tail *node
}

func New() LinkedList {
	return LinkedList{size: 0, head: nil, tail: nil}
}

// Size returns the Length of the provided LinkedList.
func (l *LinkedList) Size() int {
	return l.size
}

// PushBack adds the supplied node to the end of a LinkedList.
func (l *LinkedList) PushBack(v interface{}) {
	n := &node{data: v}

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
	n := &node{data: data}

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

// Delete removes the node with the provided key from a linkedlist.
func (l *LinkedList) Delete(position int) error {

	if l.size == 0 {
		return errors.New("Cannot delete node from empty list")
	}
	if position >= l.size {
		return errors.New("Possition outside of range.")
	}

	switch position {
	case 0:
		l.head = l.head.next
		l.size--
		if l.head == nil {
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

// Remove will removed the first occurrence of the provided data from the list, if present.
func (l *LinkedList) Remove(data interface{}) (ok bool, err error) {
	// TODO: Implimentation needed.
	return
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

// PeekFront retrieves, but does not delete the first node of a linked list.
func (l *LinkedList) PeekFront() (interface{}, error) {
	if l.head == nil {
		return 0, fmt.Errorf("Cannot find Front value in an empty linked list")
	}
	return l.head.data, nil
}

// PeekBack retrieves, but does not delete the last node of a linked list.
func (l *LinkedList) PeekBack() (interface{}, error) {
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
