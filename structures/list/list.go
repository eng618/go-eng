// Package list is a simple implementation of the linked list data structure.
package list

import (
	"fmt"
	"log"
)

// Node is the structure that makes up a single node within a linked list.
type Node struct {
	Data int
	Next *Node
}

// LinkedList is the structure of a linked list.
type LinkedList struct {
	length int
	head   *Node
	tail   *Node
}

// Len returns the length of the provided LinkedList.
func (l *LinkedList) Len() int {
	return l.length
}

// PushBack adds the supplied node to the end of a LinkedList.
func (l *LinkedList) PushBack(n *Node) {
	if l.head == nil {
		l.head = n
		l.tail = n
		l.length++
	} else {
		l.tail.Next = n
		l.tail = n
		l.length++
	}
}

// PushFront adds the supplied node to the beginning of a LinkedList.
func (l *LinkedList) PushFront(n *Node) {
	if l.head == nil {
		l.head = n
		l.tail = n
		l.length++
	} else {
		n.Next = l.head
		l.head = n
		l.length++
	}
}

// Delete removes the node with the provided key from a linkedlist.
func (l *LinkedList) Delete(key int) {

	if l.length == 0 {
		log.Println("Attempted to delete empty list")
		return
	}

	if l.head.Data == key {
		l.head = l.head.Next
		l.length--
		return
	}
	var prev *Node = nil
	n := l.head
	for n != nil && n.Data != key {
		prev = n
		n = n.Next
	}
	if n == nil {
		log.Println("Key Not found")
		return
	}
	prev.Next = n.Next
	l.length--
	log.Println("Node Deleted")
}

// Display is a helper to print a visual representation of the linked list to the console.
func (l LinkedList) Display() {
	for l.head != nil {
		if l.head.Next == nil {
			fmt.Printf("%v ->", l.head.Data)
		} else {
			fmt.Printf("%v -> ", l.head.Data)
		}
		l.head = l.head.Next
	}
	fmt.Println()
}

// Front returns the front (or first) node of a linked list.
func (l LinkedList) Front() (int, error) {
	if l.head == nil {
		return 0, fmt.Errorf("Cannot Find Front Value in an Empty linked list")
	}
	return l.head.Data, nil
}

// Back returns the back (or last) node of a linked list.
func (l LinkedList) Back() (int, error) {
	if l.head == nil {
		return 0, fmt.Errorf("Cannot Find Front Value in an Empty linked list")
	}
	return l.tail.Data, nil
}

// Reverse takes the linked list and reverses all off the values within it.
func (l *LinkedList) Reverse() {
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
