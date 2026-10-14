# Add a Package

How to add a new algorithm or data structure package so it matches repo conventions. Follow this checklist in order; `task validate` at the end must pass.

## 1. Create the package with a `doc.go`

Every package has a `doc.go` holding the top-level package documentation. Create `algo/<name>/doc.go` (or `ds/<name>/doc.go`) first, following the existing pattern — what the package provides, complexity notes, and a usage example. See `algo/sorting/doc.go` or `ds/tree/doc.go` for the shape.

## 2. Implement the API

Keep the API small and teachable. One focused file (e.g. `bubble.go`, `hash.go`) beats a sprawling one.

## 3. Add the BET trio in `*_test.go`

Every public API needs Benchmark + Example + Test:

- **Test**: external package (`package foo_test`, never the same package — except white-box tests, which live in `*_internal_test.go`, the only `testpackage` exemption). Table-driven with `t.Run(tt.name, ...)`; use `t.Helper()` in helpers (`thelper` lint).
- **Example**: an `Example*` func that teaches usage, verified by `go test` via an `// Output:` comment (`testableexamples` lint).
- **Benchmark**: a `Benchmark*` func, even when performance is not the point.

## 4. Match the style rules

- **Imports** (`gci`, `custom-order: true`): stdlib, blank line, third-party, blank line, `github.com/eng618/go-eng`. Keep module paths dotted — gofumpt sorts dotless imports as stdlib and fights `gci` grouping.
- **Formatters**: `gofumpt` (with `extra` group-params/clothe-returns/balance-calls), `gci`, `goimports`. No `gofmt`, no `golines` — line length is `lll` at 120, report-only.
- **Range style**: prefer `for range b.N` / `for range n` over classic three-clause loops where the linters suggest it (`intrange`, `modernize`).

## 5. Wire up docs

- Add the package to the [package index](../reference/packages.md) and the README "What's Inside" table.
- If it resolves a planned item, tick it off in the [inventory](../project/inventory.md).

## 6. Validate

```sh
task validate
```

Fix with `task lint-fix` where the fixes are mechanical, then re-run. See [Run tasks](run-tasks.md) for targeted runs.
