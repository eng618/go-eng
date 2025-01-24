/*
Package list demonstrates a linked list data structure.

# Linked List Data Structure

This package provides a thread-safe implementation of a singly linked list with head and tail pointers. The head and tail pointers allow for efficient O(1) time complexity for both adding and removing elements from the ends of the list.

## Features

- Add elements to the front or back of the list.
- Remove elements from the list.
- Peek at the front or back elements without removing them.
- Reverse the list in place.
- Display the list elements.
- Concurrent access control using a mutex.

## Example Usage

```go
package main

import (

	"fmt"
	"github.com/eng618/go-eng/ds/list"

)

	func main() {
		ll := list.New()
		ll.PushBack(1)
		ll.PushBack(2)
		ll.PushFront(0)
		ll.Display()
		// Output: 0 -> 1 -> 2 ->
	}

```

I started with this [code source](https://divyanshushekhar.com/golang-linked-list/) and expanded on it to test and document the code using Go's best practices. Also, while writing tests, I found a bug in the code that it was not properly deleting the final element of a list.
*/
package list
