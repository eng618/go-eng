# Commands

Lookup for `Taskfile.yml` targets. `Taskfile.yml` is the source of truth — this page mirrors it and the [Run tasks](../how-to/run-tasks.md) guide explains when to use each.

## Aggregate (all modules)

| Target          | Effect                                                                                                     |
| --------------- | ---------------------------------------------------------------------------------------------------------- |
| `task validate` | `lint` → `test` → `build` (mirrors CI)                                                                     |
| `task lint`     | Lint all modules                                                                                           |
| `task lint-fix` | Lint-fix all modules                                                                                       |
| `task test`     | Test all modules (`CGO_ENABLED=1 go test ./... -cover -race -coverprofile=coverage.out -covermode=atomic`) |
| `task build`    | Build all modules                                                                                          |
| `task bench`    | Benchmark all modules                                                                                      |
| `task tidy`     | `go mod tidy` everywhere                                                                                   |

## Per-module

Prefixes: `root`, `circuit-breaker`, `master-slave`, `contexts`, `generics`, `web-service-gin`. Suffixes: `-build`, `-lint`, `-lint-fix`, `-test`, `-bench`, `-tidy`, `-update`. Plus `root`, `generics`, `web-service-gin`, `circuit-breaker`, `master-slave` run targets. Examples:

| Target                       | Effect                          |
| ---------------------------- | ------------------------------- |
| `task root-test`             | Test root module only           |
| `task circuit-breaker-test`  | Test the circuit-breaker module |
| `task web-service-gin-build` | Build the gin example           |

## Raw `go` (from the module directory)

```sh
go test ./algo/sorting/... -run TestBubbleSort
go test ./algo/sorting/... -bench=BenchmarkBubbleSort
```
