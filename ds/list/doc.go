/*
Package list demonstrates linked list data structures.

# Linked List Data Structures

This package provides two implementations of linked lists:

1. A custom implementation of a singly linked list with head and tail pointers.
2. A wrapper around the standard library's container/list package.

## Features

- Add elements to the front or back of the list.
- Remove elements from the list.
- Peek at the front or back elements without removing them.
- Reverse the list in place (custom implementation).
- Display the list elements.
- Concurrent access control using a mutex (custom implementation).

## Example Usage (Custom Implementation)

```go
package main

import (

	"fmt"
	"github.com/eng618/go-eng/ds/list"

)

	func main() {
		ll := list.NewLinkedList()
		ll.PushBack(1)
		ll.PushBack(2)
		ll.PushFront(0)
		ll.Display()
		// Output: 0 -> 1 -> 2 ->
	}

```

## Example Usage (Standard Library Wrapper)

```go
package main

import (

	"fmt"
	"github.com/eng618/go-eng/ds/list"

)

	func main() {
		l := list.NewList()
		l.PushBack(1)
		l.PushFront(0)
		l.PushBack(2)
		l.Iterate(func(value interface{}) {
			fmt.Println(value)
		})
		// Output:
		// 0
		// 1
		// 2
	}

```

I started with this [code source](https://divyanshushekhar.com/golang-linked-list/)
and expanded on it to test and document the code using Go's best practices.
Also, while writing tests, I found a bug in the code that it was not properly
deleting the final element of a list.
*/
package list
