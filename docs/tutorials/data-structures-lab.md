# Data Structures Lab

Learn the core data structures by using them. By the end you will have pushed and popped a stack, enqueued and dequeued from both queue variants, and built and reversed a linked list. Allow about 25 minutes.

## What you need

- The [Getting Started](getting-started.md) tour completed.
- A scratch directory outside the repo (e.g. `/tmp/dslab`) so experiments never get committed.

## 1. Stack: last in, first out

```go
package main

import (
	"fmt"

	"github.com/eng618/go-eng/ds/stack"
)

func main() {
	s := stack.New()
	s.Push(25)
	s.Push(1)
	s.Push(2)

	v, err := s.Pop()
	fmt.Println("Popped:", v, "err:", err)
}
```

Expected output: `Popped: 2 err: <nil>` — the last item pushed is the first out. The stack is safe for concurrent use via a read-write mutex; `NewSeeded` builds one pre-filled.

Check the teaching trio in `ds/stack/stack_test.go` (Example, table-driven Test, Benchmark), then run it:

```sh
go test ./ds/stack/... -v
```

## 2. Queue: first in, first out

Two implementations share the same idea. Extend your scratch program:

```go
import "github.com/eng618/go-eng/ds/queue"

sq := queue.SliceQueue{}
sq.Enqueue("a")
sq.Enqueue("b")
front, _ := sq.Dequeue()
fmt.Println("Slice queue dequeued:", front)

lq := queue.LinkedQueue{}
lq.Enqueue("a")
lq.Enqueue("b")
front, _ = lq.Dequeue()
fmt.Println("Linked queue dequeued:", front)
```

Expected outcome: both print `a` — queues release the oldest item first. (`Peek` looks without removing; `Length` counts items.) A circular, dynamically-growing array variant is planned — see the [inventory](../project/inventory.md).

## 3. Linked list: build and reverse

```go
import "github.com/eng618/go-eng/ds/linkedlist"

ll := linkedlist.NewLinkedList()
ll.PushBack(20)
ll.PushBack(30)
ll.PushFront(70)
fmt.Println("Length =", ll.Length())
ll.Display()

ll.Reverse()
ll.Display()
front, _ := ll.PeekFront()
back, _ := ll.PeekBack()
fmt.Println("Front =", front, "Back =", back)
```

Expected outcome: length 3, the list displayed, then displayed reversed, with front/back swapped. The custom list carries head and tail pointers plus a mutex; `ds/linkedlist` also ships a `container/list` wrapper (`NewList`) if you prefer the standard library.

## 4. Binary search tree: ordered storage

```go
import "github.com/eng618/go-eng/ds/tree"

tr := tree.New()
for _, v := range []int{10, 5, 15, 3, 7} {
	tr.Insert(v)
}
fmt.Println("Found 7:", tr.Search(7))
fmt.Println("Found 100:", tr.Search(100))
```

Expected outcome: `true`, then `false`. Values left of a node are smaller, right are larger or equal — the BST property that makes search logarithmic on balanced trees.

## Where next

- [Big-O cheatsheet](../explanation/big-o.md) — name the complexity of what you just ran.
- [Add a package](../how-to/add-a-package.md) — contribute the next structure using the BET checklist.
