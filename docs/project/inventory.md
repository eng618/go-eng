# Implementation Inventory

Sources: former `README.md` "Still to come" section (since removed; README now links here), open/closed GitHub issues (`gh issue list`), disk scan of `algo/`, `ds/`, `design-patterns/`, `lib/`, `examples/`.
Out of scope for this doc: infra issues, `interview/`, `leet/`.

## Algorithms — Implemented

| Package          | API                                                           |
| ---------------- | ------------------------------------------------------------- |
| `algo/anagrams`  | `IsAnagram`                                                   |
| `algo/fibonacci` | `Basic`, `BottomUp`, `Dynamic`                                |
| `algo/search`    | `Binary` (loop + recursion), `Linear`                         |
| `algo/sorting`   | `BubbleSort`, `MergeSort`, `Merge`, `MergeSimple`, `Standard` |
| `algo/vowels`    | `Count`                                                       |

## Algorithms — Outstanding

- [ ] Sorting: Insertion (`README.md`, issue #8)
- [ ] Sorting: Quick sort, recursive (`README.md`, issue #6)
- [ ] Search: Breadth-first search / BFS (issue #55; builds on #54)
- [ ] Search: Depth-first search / DFS (issue #56; builds on #54)
- [ ] Recursion topic (`README.md` — vague; `fibonacci` already covers recursion, needs scoping)

## Data Structures — Implemented

| Package         | API                                                      |
| --------------- | -------------------------------------------------------- |
| `ds/linkedlist` | Custom singly linked list + `container/list` wrapper     |
| `ds/queue`      | `SliceQueue`, `LinkedListQueue`                          |
| `ds/stack`      | Slice-based `Stack` (`New`, `NewSeeded`, `Push`, `Pop`)  |
| `ds/hashtable`  | `Table` (`New`, `Set`, `Get`, `Keys`, `Values`, `Print`) |
| `ds/tree`       | BST (`New`, `NewWithRoot`, `Insert`, `Search`, `Delete`) |

## Data Structures — Outstanding

- [ ] Double linked list (issue #4; `README.md` prematurely marks `[x]`)
- [ ] Circular dynamic array queue (issue #3 body; slice + linked-list variants done, circular variant missing)
- [ ] Graph (issue #54; unblocks BFS/DFS)
- [ ] `container/heap` wrapper (issue #57)
- [ ] `container/ring` wrapper (issue #58)

## Design Patterns / Libs / Examples — Implemented

| Area    | Package                                    | Notes                                     |
| ------- | ------------------------------------------ | ----------------------------------------- |
| Pattern | `design-patterns/circuit-breaker/breaker`  | Closed/Open/Half-Open, `Execute`          |
| Pattern | `design-patterns/master-slave/masterslave` | Task distribution, health checks, metrics |
| Lib     | `lib/file` (`engfile`)                     | `ToFile`                                  |
| Example | `examples/contexts`                        | Context timeout                           |
| Example | `examples/generics`                        | Go generics tutorial                      |
| Example | `examples/web-service-gin`                 | Gin albums API                            |

No outstanding design-pattern, lib, or example items are promised in docs or issues.

## Candidate additions (proposed, not promised)

High-value gaps that fit the teaching model (BET) and complement what exists. Not sourced from issues — add to build order only if accepted.

### Algorithms

- [ ] Selection sort — completes the elementary-sort trio with bubble/insertion; trivial to teach.
- [ ] Heap sort — natural follow-on to `container/heap` + merge/quick; shows `O(n log n)` in place.
- [ ] Sliding window / two pointers — high interview value; pairs with `anagrams`/`vowels` string work.
- [ ] String matching (KMP or Rabin-Karp) — next step after anagram/vowel basics.
- [ ] Dijkstra's shortest path — pairs with planned Graph + heap; gives BFS/DFS a payoff.
- [ ] Binary search bounds (`LowerBound`/`UpperBound`) — small extension of existing `search`, frequently asked.

### Data Structures

- [ ] Set (map-based) — idiomatic Go, conspicuously missing.
- [ ] Priority queue via `container/heap` — the canonical heap demo; unblocks Dijkstra.
- [ ] LRU cache — classic interview item; ties together map + doubly-linked list (#4).
- [ ] Trie — string DS that extends the anagrams/vowels thread.
- [ ] Union-Find / Disjoint set — small, powerful companion to Graph.

### Design Patterns (Go-idiomatic first)

- [ ] Functional options — the most idiomatic Go pattern; high reuse value.
- [ ] Worker pool / pipeline / fan-in-fan-out — core Go concurrency patterns; pairs with master-slave + contexts.
- [ ] Retry with backoff + rate limiter — natural companions to circuit-breaker.
- [ ] Strategy (interfaces), Decorator/middleware (`http.Handler`), Observer/pub-sub — small, classic GoF patterns that teach interfaces well.

## Discrepancies (README vs issues vs disk)

- Bubble sort: code existed with issue #7 open → verified and closed #7 (BET present, tests pass).
- Hash table: code existed with issue #12 open → verified and closed #12 (BET present, tests pass).
- BST: code exists and issue #10 closed; stale `README.md` checklist removed in README redesign → resolved.
- Double linked list: old `README.md` listed `[x]` with issue #4 open and no code → fixed in README redesign (no longer claimed); #4 stays open.
- Queue: old `README.md` listed `[x]` with issue #3 (circular array) open → README redesign now claims only slice + linked-list variants; #3 stays open.

## Suggested build order

1. Insertion sort (#8)
2. Quick sort (#6)
3. Circular queue (#3)
4. Double linked list (#4)
5. `container/heap`
6. Graph + BFS/DFS
7. `container/ring`
8. Scope "Recursion" (decide if `fibonacci` suffices)
