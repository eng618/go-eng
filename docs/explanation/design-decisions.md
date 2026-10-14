# Design Decisions

Why this repo is shaped the way it is. This page explains the reasoning; for instructions see the [how-to guides](../how-to/run-tasks.md).

## A workspace, not one module

The root module holds the stable teaching packages (`algo/`, `ds/`, `lib/`). Anything with its own dependencies or runnable server — the two design-pattern demos and the three examples — is an independent module under `go.work`. This keeps the root dependency-free while letting each demo evolve on its own. The cost is that `go` commands only see one module at a time, which is why `Taskfile.yml` provides `task <module>-<action>` cross-module targets.

## Every public API is taught (BET)

Each public API ships three companions: a Benchmark, an Example with `// Output:`, and a table-driven Test. Examples are verified by `go test`, so documentation cannot drift from behavior silently. Benchmarks exist even where performance is irrelevant — they prove the API holds up under repetition. External `foo_test` packages keep tests honest about the public surface; `*_internal_test.go` is reserved for the rare white-box case.

## `doc.go` owns package documentation

Top-level package docs live in `doc.go`, not in `docs/`. Prose in `docs/` links to packages; it never duplicates them. That is also why the [package index](../reference/packages.md) points at `doc.go` files and pkg.go.dev rather than restating APIs.

## `Taskfile.yml` is the source of truth

All workflows — lint, test, build, bench, tidy — are Task targets, and CI runs the same `task validate` sequence a contributor runs locally. Raw commands appear in docs only for single-test/single-benchmark runs where a Task target would add nothing.

## Strict, teaching-oriented linting

The lint set (`gofumpt`, `gci` with a three-group import order, `modernize`, `testableexamples`, `testpackage`, and friends) encodes the conventions: dotted module paths so import grouping is stable, `for range n` idioms, helper-aware tests. The `env -u GOROOT -u GOTOOLDIR` wrapper and `GOTOOLCHAIN=local` pin exist because the environment once shipped a stale toolchain — see [Troubleshoot the toolchain](../how-to/troubleshoot-toolchain.md).
