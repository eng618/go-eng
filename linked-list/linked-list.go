// package linkedlist is a simple implimentation of the linked list data structure.
package linkedlist

import "fmt"

// LLNode is the structure that makes up a single node withing a linked list.
type LLNode struct {
	Data int
	Next *LLNode
}

// LinkedList is the structure of a linked list.
type LinkedList struct {
	length int
	head   *LLNode
	tail   *LLNode
}

// Len returns the length of the provided LinkedList.
func (l *LinkedList) Len() int {
	return l.length
}

// PushBack adds the supplied node to the end of a LinkedList.
func (l *LinkedList) PushBack(n *LLNode) {
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
func (l *LinkedList) PushFront(n *LLNode) {
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

	if l.head.Data == key {
		l.head = l.head.Next
		l.length--
		return
	}
	var prev *LLNode = nil
	n := l.head
	for n != nil && n.Data != key {
		prev = n
		n = n.Next
	}
	if n == nil {
		fmt.Println("Key Not found")
		return
	}
	prev.Next = n.Next
	l.length--
	fmt.Println("Node Deleted")
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

// Reverse takes the linked list and reverses all off the values withing it.
func (l *LinkedList) Reverse() {
	n := l.head
	l.tail = l.head
	var prev *LLNode
	for n != nil {
		temp := n.Next
		n.Next = prev
		prev = n
		n = temp
	}
	l.head = prev
}
