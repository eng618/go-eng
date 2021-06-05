package main

import (
	"fmt"
	"log"

	"github.com/ENG618/go-eng/search/binary"
	"github.com/ENG618/go-eng/structures/list"
	"github.com/ENG618/go-eng/structures/stack"
)

func main() {
	fmt.Println("go-eng")
	showLinkedList()
	showBinarySearch()
	showStack()
}

func showLinkedList() {
	log.Println("showLinkedList entered")
	fmt.Println("Below are example outputs of the list package in action")
	ll := list.LinkedList{}
	node1 := &list.Node{Data: 20}
	node2 := &list.Node{Data: 30}
	node3 := &list.Node{Data: 40}
	node4 := &list.Node{Data: 50}
	node5 := &list.Node{Data: 70}
	ll.PushBack(node1)
	ll.PushBack(node2)
	ll.PushBack(node3)
	ll.PushFront(node4)
	ll.PushFront(node5)
	fmt.Println("Length = ", ll.Len())
	ll.Display()
	ll.Delete(40)
	ll.Reverse()
	fmt.Println("Length = ", ll.Len())
	ll.Display()
	front, _ := ll.Front()
	back, _ := ll.Back()
	fmt.Println("Front = ", front)
	fmt.Println("Back = ", back)

	// Output:
	// Length =  5
	// 20 -> 30 -> 40 -> 50 -> 70 ->
	// Node Deleted
	// Length =  4
	// 70 -> 50 -> 30 -> 20 ->
	// Front =  70
	// Back =  20
	fmt.Println("-----END list-----")
}

func showBinarySearch() {
	log.Println("showBinarySearch entered")
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	if v, ok := binary.BinarySearchForLoop(xi, 5); ok {
		fmt.Println("Found 5 at index", v)
	}

	if v, ok := binary.BinarySearchForLoop(xi, 25); ok {
		fmt.Println("Found 25 at index", v)
	} else {
		fmt.Println("target number no found in slice")
	}
}

func showStack() {
	log.Println("showStack entered")
	s := stack.Stack{}
	s.Initialize()

	s.Push(25)
	s.Push(1)
	s.Push(2)
	if v, ok := s.Pop(); ok {
		fmt.Println("Pop returned", v)
	}
}
