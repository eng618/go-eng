# Copilot Instructions

## Repository Overview

This is a Go learning/exploration repository (`github.com/eng618/go-eng`) containing algorithm implementations, data structures, design patterns, and example applications. It is organized as a **Go workspace** (`go.work`) with multiple independent modules.

### Module Layout

| Path                              | Module                                            |
| --------------------------------- | ------------------------------------------------- |
| `.`                               | `github.com/eng618/go-eng` (root: algos, DS, lib) |
| `design-patterns/circuit-breaker` | Circuit Breaker pattern                           |
| `design-patterns/master-slave`    | Master-Slave pattern                              |
| `examples/contexts`               | Context usage examples                            |
| `examples/generics`               | Generics examples                                 |
| `examples/web-service-gin`        | REST API with Gin                                 |

## Build, Test, and Lint

The preferred task runner is [Task](https://taskfile.dev) (`task`). `make` targets also exist with equivalent behavior.

### All modules

```sh
task build      # build all modules
task test       # test all modules (race detection + coverage)
task lint       # lint all modules
task lint-fix   # lint with auto-fix
task validate   # lint + test + build
```

### Root module only

```sh
task root-test
# equivalent:
CGO_ENABLED=1 go test ./... -cover -race -coverprofile=coverage.out -covermode=atomic
```

### Single test

```sh
go test ./algo/sorting/... -run TestBubbleSort
```

### Sub-module tasks follow the pattern `task <module>-<action>`:

```sh
task circuit-breaker-test
task master-slave-lint-fix
task web-service-gin-build
```

### Benchmarks

```sh
task bench      # all modules
task root-bench # root only
go test ./algo/sorting/... -bench=BenchmarkBubbleSort
```

## Key Conventions

### Tool Versions

`.tool-versions` (managed by `asdf`/`mise`) is the source of truth for tool versions. When upgrading `golangci-lint` or `golang`, update **both** `.tool-versions` and the matching `version:` field in `.github/workflows/ci.yml` in the same commit to avoid drift.

### Go Version

Root `go.mod` and the workspace (`go.work`) both target **Go 1.26**.

### Import Ordering (enforced by `gci` + `golangci-lint`)

Three groups in this order, separated by blank lines:

1. Standard library
2. Third-party packages
3. Local (`github.com/eng618/go-eng`)

### Formatting

- **gofumpt** (stricter than gofmt) with `extra-rules: true`
- **golines** enforces a max line length of **120 characters**

### Test Patterns

- Use **external test packages** (`package foo_test`) — not `package foo`
- Use **table-driven tests** with `t.Run(tt.name, ...)` subtests
- Include `Example*` functions for public APIs (they are verified by `go test`)
- Include `Benchmark*` functions for performance-sensitive code

### Package Documentation

Each package has a `doc.go` file with a package-level doc comment. New packages should follow this pattern.

### Changelog

Use `git-chglog` to update `CHANGELOG.md`:

```sh
git-chglog -o CHANGELOG.md
```

Conventional commits are expected (e.g., `feat:`, `fix:`, `chore:`).
