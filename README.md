# go-eng

[![build-test-lint](https://github.com/eng618/go-eng/actions/workflows/ci.yml/badge.svg)](https://github.com/eng618/go-eng/actions/workflows/ci.yml)
[![release](https://github.com/eng618/go-eng/actions/workflows/release.yml/badge.svg)](https://github.com/eng618/go-eng/actions/workflows/release.yml)
[![codecov](https://codecov.io/gh/eng618/go-eng/branch/main/graph/badge.svg?token=L1D7JSRK98)](https://codecov.io/gh/eng618/go-eng)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/eng618/go-eng)](https://pkg.go.dev/github.com/eng618/go-eng#section-directories)
[![Go Version](https://img.shields.io/badge/go-1.27-blue)](https://go.dev)

A hands-on Go workspace for learning algorithms, data structures, design patterns, and idiomatic Go. Every public API is taught, not just shipped: documented in `doc.go` and demonstrated with Benchmark + Example + Test.

## Contents

- [About](#about)
- [What's Inside](#whats-inside)
- [Project Structure](#project-structure)
- [Getting Started](#getting-started)
- [Usage](#usage)
- [Documentation](#documentation)
- [Development](#development)
- [Testing Conventions](#testing-conventions)
- [Roadmap](#roadmap)
- [Contributing](#contributing)
- [License](#license)

## About

`go-eng` is a personal learning repository that doubles as a reusable library of small, well-tested Go packages. The root `main.go` is a demo only — the real value is in the subpackages under `algo/`, `ds/`, `lib/`, plus the isolated modules under `design-patterns/` and `examples/`.

Full API docs live on [pkg.go.dev](https://pkg.go.dev/github.com/eng618/go-eng#section-directories).

## What's Inside

### Algorithms

| Package | Provides |
| --- | --- |
| [`algo/anagrams`](algo/anagrams/doc.go) | Anagram detection (`IsAnagram`) |
| [`algo/fibonacci`](algo/fibonacci/doc.go) | Recursive, iterative, and memoized Fibonacci |
| [`algo/search`](algo/search/doc.go) | Binary search (loop + recursion), linear search |
| [`algo/sorting`](algo/sorting/doc.go) | Bubble sort, merge sort, std-lib merge helpers |
| [`algo/vowels`](algo/vowels/vowels.go) | Vowel counting |

### Data Structures

| Package | Provides |
| --- | --- |
| [`ds/linkedlist`](ds/linkedlist/doc.go) | Custom singly linked list + `container/list` wrapper |
| [`ds/queue`](ds/queue/doc.go) | Slice-backed and linked-list-backed queues |
| [`ds/stack`](ds/stack/doc.go) | Thread-safe slice-backed stack |
| [`ds/hashtable`](ds/hashtable/doc.go) | Basic hash table (`Set`, `Get`, `Keys`, `Values`) |
| [`ds/tree`](ds/tree/doc.go) | Binary search tree (`Insert`, `Search`, `Delete`) |

### Design Patterns, Libs, and Examples

| Area | Package | Notes |
| --- | --- | --- |
| Pattern | `design-patterns/circuit-breaker/breaker` | Closed / Open / Half-Open states, thread-safe |
| Pattern | `design-patterns/master-slave/masterslave` | Concurrent task distribution, health checks, metrics |
| Lib | `lib/file` | `ToFile` file-writing helper |
| Example | `examples/contexts` | Context timeout pattern |
| Example | `examples/generics` | Generics tutorial code |
| Example | `examples/web-service-gin` | Gin albums REST API tutorial |

See [docs/project/inventory.md](docs/project/inventory.md) for the authoritative implemented-vs-planned list.

## Project Structure

Go workspace (`go.work`) with one root module plus five independent modules:

```text
.
├── algo/                  # sorting, searching, fibonacci, anagrams, vowels
├── ds/                    # linkedlist, queue, stack, hashtable, tree
├── lib/                   # small helpers (file)
├── main.go                # demo entrypoint (not a library)
├── docs/                  # tutorials, how-to, reference, explanation, project
├── design-patterns/
│   ├── circuit-breaker/   # separate module
│   └── master-slave/      # separate module
└── examples/
    ├── contexts/          # separate module
    ├── generics/          # separate module
    └── web-service-gin/   # separate module
```

| Module | Path |
| --- | --- |
| Root | `github.com/eng618/go-eng` |
| Circuit breaker | `design-patterns/circuit-breaker` |
| Master-slave | `design-patterns/master-slave` |
| Contexts | `examples/contexts` |
| Generics | `examples/generics` |
| Web service | `examples/web-service-gin` |

`go` commands for a submodule must run in that directory. The cross-module equivalent is `task <module>-<action>`, e.g. `task circuit-breaker-test`.

## Getting Started

### Prerequisites

- Go 1.27 (see `.tool-versions`; `mise`/`asdf` recommended)
- [Task](https://taskfile.dev) 3.x
- [golangci-lint](https://golangci-lint.run) 2.x (invoked via `task lint`)

### Install

```sh
git clone https://github.com/eng618/go-eng.git
cd go-eng
go work sync
```

### Run the demo

```sh
go run main.go
```

Run an isolated module:

```sh
task generics          # cd examples/generics && go run main.go
task circuit-breaker   # cd design-patterns/circuit-breaker && go run main.go
```

## Usage

Import any root-module package by its full path:

```go
import (
    "fmt"

    "github.com/eng618/go-eng/algo/search"
    "github.com/eng618/go-eng/algo/sorting"
    "github.com/eng618/go-eng/ds/stack"
)

func main() {
    fmt.Println(sorting.MergeSort([]int{5, 4, 3, 2, 1}))

    idx, ok := search.BinaryLoop([]int{1, 2, 3, 4, 5}, 4)
    fmt.Println(idx, ok)

    s := stack.New()
    s.Push(1)
    v, _ := s.Pop()
    fmt.Println(v)
}
```

Circuit breaker (separate module — see its `go.mod`):

```go
cb := breaker.NewCircuitBreaker(5, 10*time.Second)

err := cb.Execute(func() error {
    return someRiskyOperation()
})
```

## Documentation

- Start with the [documentation index](docs/index.md) — tutorials, how-to guides, reference, and explanation.
- API reference: [pkg.go.dev](https://pkg.go.dev/github.com/eng618/go-eng#section-directories)
- [Big-O cheatsheet](docs/explanation/big-o.md) — complexity rules used across the labs.
- [Implementation inventory](docs/project/inventory.md) — what is implemented vs. planned.
- Each package documents itself in `doc.go` with runnable examples

## Development

`Taskfile.yml` is the source of truth for all workflows. CI runs `task lint`, then `task test`, then `task build`.

```sh
task validate   # lint -> test -> build (mirrors CI)
task lint       # all modules; fix with task lint-fix
task test       # all modules, with -race and coverage
task build      # all modules
task bench      # benchmarks, all modules
```

Targeted runs (from the module directory):

```sh
task root-test
go test ./algo/sorting/... -run TestBubbleSort
go test ./algo/sorting/... -bench=BenchmarkBubbleSort
```

Notes:

- Do not run bare `golangci-lint run ./...`; `Taskfile` wraps it as `env -u GOROOT -u GOTOOLDIR golangci-lint run ./...`.
- `Taskfile` sets `GOTOOLCHAIN=local`.
- Generated, gitignored artifacts: `coverage.out`, `linter-output.json` — do not commit.

## Testing Conventions

Every public API follows the BET method:

- **Benchmark** — a `Benchmark*` function, even when perf is not the point.
- **Example** — an `Example*` function with `// Output:` that teaches usage.
- **Test** — table-driven, external `package foo_test`, using `t.Run(tt.name, ...)`.

White-box tests live in `*_internal_test.go` (the only `testpackage` exemption). New packages must add a `doc.go`.

Import grouping (`gci`): stdlib, blank line, third-party, blank line, `github.com/eng618/go-eng`. Line length is 120 (`lll`, report-only).

## Roadmap

Planned work (insertion sort, quick sort, BFS/DFS, graph, `container/heap`, `container/ring`, circular queue, doubly linked list) is tracked in [docs/project/inventory.md](docs/project/inventory.md), sourced from open issues and the former "Still to come" list. Work items are intentionally small so they can be done one at a time.

## Contributing

Conventional commits (`feat:`, `fix:`, `chore:`, …) are required — `release-please` owns `CHANGELOG.md`, so never edit it by hand. Run `task validate` before pushing.

## License

See [LICENSE](LICENSE).
