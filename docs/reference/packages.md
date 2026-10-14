# Packages

Index of root-module packages. Each package documents itself in its `doc.go`; full API detail lives on [pkg.go.dev](https://pkg.go.dev/github.com/eng618/go-eng#section-directories).

## Algorithms

| Package          | Provides                                                      | `doc.go`                                                       |
| ---------------- | ------------------------------------------------------------- | -------------------------------------------------------------- |
| `algo/anagrams`  | `IsAnagram`                                                   | [doc.go](../../algo/anagrams/doc.go)                           |
| `algo/fibonacci` | `Basic`, `BottomUp`, `Dynamic`                                | [doc.go](../../algo/fibonacci/doc.go)                          |
| `algo/search`    | Binary (loop + recursion), linear                             | [doc.go](../../algo/search/doc.go)                             |
| `algo/sorting`   | `BubbleSort`, `MergeSort`, `Merge`, `MergeSimple`, `Standard` | [doc.go](../../algo/sorting/doc.go)                            |
| `algo/vowels`    | `Count`                                                       | [source](../../algo/vowels/vowels.go) (`doc.go` not yet added) |

## Data structures

| Package         | Provides                                                 | `doc.go`                             |
| --------------- | -------------------------------------------------------- | ------------------------------------ |
| `ds/linkedlist` | Custom singly linked list + `container/list` wrapper     | [doc.go](../../ds/linkedlist/doc.go) |
| `ds/queue`      | `SliceQueue`, `LinkedQueue`                              | [doc.go](../../ds/queue/doc.go)      |
| `ds/stack`      | `Stack` (`New`, `NewSeeded`, `Push`, `Pop`)              | [doc.go](../../ds/stack/doc.go)      |
| `ds/hashtable`  | `Table` (`New`, `Set`, `Get`, `Keys`, `Values`, `Print`) | [doc.go](../../ds/hashtable/doc.go)  |
| `ds/tree`       | BST (`New`, `NewWithRoot`, `Insert`, `Search`, `Delete`) | [doc.go](../../ds/tree/doc.go)       |

## Libs

| Package                | Provides |
| ---------------------- | -------- |
| `lib/file` (`engfile`) | `ToFile` |

## Separate modules

| Package                                    | `doc.go`                                                        |
| ------------------------------------------ | --------------------------------------------------------------- |
| `design-patterns/circuit-breaker/breaker`  | [doc.go](../../design-patterns/circuit-breaker/breaker/doc.go)  |
| `design-patterns/master-slave/masterslave` | [doc.go](../../design-patterns/master-slave/masterslave/doc.go) |
| `examples/generics`                        | [doc.go](../../examples/generics/doc.go)                        |
| `examples/web-service-gin`                 | [doc.go](../../examples/web-service-gin/doc.go)                 |
