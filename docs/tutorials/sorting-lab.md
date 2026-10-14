# Sorting Lab

Learn the sorting packages by sorting real slices. By the end you will have used `BubbleSort` and `MergeSort`, read their Example tests, and compared their complexity. Allow about 20 minutes.

## What you need

- The [Getting Started](getting-started.md) tour completed (repo cloned, demo runs).
- The [Big-O cheatsheet](../explanation/big-o.md) open in a second tab — we will refer to it.

## 1. Bubble sort: sort your first slice

Create a scratch file outside the repo (so it never gets committed), for example `/tmp/sortlab/main.go`:

```go
package main

import (
	"fmt"

	"github.com/eng618/go-eng/algo/sorting"
)

func main() {
	data := []int{5, 3, 8, 4, 2}
	fmt.Println(sorting.BubbleSort(data))
}
```

Run it with `go run` from the repo root so the module resolves:

```sh
go run /tmp/sortlab/main.go
```

Expected output:

```text
[2 3 4 5 8]
```

Bubble sort compares neighbours and swaps them, repeating until sorted: `O(n)` on already-sorted input, `O(n²)` in the worst case.

## 2. See how it is taught

Open `algo/sorting/bubble_test.go`. Notice the three parts — this is the BET method every package in this repo follows:

1. `ExampleBubbleSort` — runnable usage with an `// Output:` comment (verified by `go test`).
2. `TestBubbleSort` — table-driven cases (`sorted`, `reverse`, `unsorted`, `empty`, …) run via `t.Run(tt.name, ...)`.
3. `BenchmarkBubbleSort` — exercises the function under repetition.

Run all three yourself:

```sh
go test ./algo/sorting/... -run 'TestBubbleSort|ExampleBubbleSort' -v
go test ./algo/sorting/... -bench=BenchmarkBubbleSort
```

## 3. Merge sort: divide and conquer

Extend your scratch program:

```go
data := []int{3, 44, 38, 5, 47, 15, 36, 26, 27, 2, 46, 4, 19, 50, 48}
fmt.Println("Before:", data)
fmt.Println("After merge sort:", sorting.MergeSort(data))

left := []int{1, 3, 5}
right := []int{2, 4, 6}
fmt.Println("Merged:", sorting.Merge(left, right))
```

Expected outcome: the shuffled slice comes back ordered, and the two halves merge into `[1 2 3 4 5 6]`. Merge sort splits the input into halves, sorts each, and merges — `O(n log n)`, the log-linear class from the [cheatsheet](../explanation/big-o.md).

## 4. Compare the two

| Algorithm    | Best         | Worst        | Idea               |
| ------------ | ------------ | ------------ | ------------------ |
| `BubbleSort` | `O(n)`       | `O(n²)`      | Neighbour swaps    |
| `MergeSort`  | `O(n log n)` | `O(n log n)` | Split, sort, merge |

That gap is why both exist here: bubble sort teaches the mechanism, merge sort teaches the better complexity class.

## Where next

- [Data Structures Lab](data-structures-lab.md) — store and order data with stacks, queues, and lists.
- [Implementation Inventory](../project/inventory.md) — insertion sort and quick sort are the planned next labs.
