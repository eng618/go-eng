# Getting Started

Your first guided tour of this repo. By the end you will have cloned the workspace, run the demo, and run your first test. Allow about 15 minutes.

## What you need

- Go 1.27 (see `.tool-versions`; `mise` or `asdf` recommended)
- [Task](https://taskfile.dev) 3.x
- A terminal at the repo root

No prior knowledge of this codebase is assumed. Familiarity with basic Go syntax helps.

## 1. Clone and sync

```sh
git clone https://github.com/eng618/go-eng.git
cd go-eng
go work sync
```

`go work sync` links the six workspace modules (root plus five independent ones) so imports resolve across them.

## 2. Run the demo

```sh
go run main.go
```

You will see four sections of output: `LinkedList`, `BinarySearch` (loop and recursion), `Stack`, and `MergeSort`. The root `main.go` is a demo only — its job is to show a few subpackages working, not to serve as a library entrypoint.

## 3. Run your first test

Tests live next to the code they teach. Run the sorting package:

```sh
go test ./algo/sorting/... -run TestBubbleSort -v
```

Expected outcome: each subtest (`sorted`, `reverse`, `unsorted`, `empty`, …) passes. Every test here follows the BET method — Benchmark, Example, Test — which you will meet properly in the labs.

## 4. Run one benchmark

```sh
go test ./algo/sorting/... -bench=BenchmarkBubbleSort
```

Expected outcome: a `BenchmarkBubbleSort` line with ns/op figures. Benchmarks exist even when performance is not the point; they prove the API is exercised under repetition.

## Where next

- [Sorting Lab](sorting-lab.md) — get hands-on with `BubbleSort` and `MergeSort`.
- [Data Structures Lab](data-structures-lab.md) — push, enqueue, and link lists yourself.
- [Big-O cheatsheet](../explanation/big-o.md) — the complexity vocabulary the labs use.
