# Run Tasks

How to run the standard workflows. `Taskfile.yml` is the source of truth — prefer `task` targets over raw `go` commands.

## Full validation (mirrors CI)

CI runs lint, then test, then build. Reproduce it locally with one command:

```sh
task validate
```

This runs `task lint`, `task test`, and `task build` in the required order.

## Everyday commands

| Goal                               | Command                                |
| ---------------------------------- | -------------------------------------- |
| Lint all modules                   | `task lint` (fix with `task lint-fix`) |
| Test all modules (race + coverage) | `task test`                            |
| Build all modules                  | `task build`                           |
| Benchmark all modules              | `task bench`                           |

## One module at a time

`go` commands for a submodule must run in that directory. The equivalent from the repo root is `task <module>-<action>`:

```sh
task root-test                # root module only
task circuit-breaker-test     # design-patterns/circuit-breaker
task web-service-gin-build    # examples/web-service-gin
```

Available module prefixes: `root`, `circuit-breaker`, `master-slave`, `contexts`, `generics`, `web-service-gin`. Each supports `-build`, `-lint`, `-lint-fix`, `-test`, `-bench`, `-tidy`, `-update`. See [Commands](../reference/commands.md) for the full lookup.

## One test or benchmark

From the module directory:

```sh
go test ./algo/sorting/... -run TestBubbleSort
go test ./algo/sorting/... -bench=BenchmarkBubbleSort
```

## If the toolchain misbehaves

A stale `GOROOT` in your shell breaks builds (see [Troubleshoot the toolchain](troubleshoot-toolchain.md)). Prefix affected commands with `env -u GOROOT -u GOTOOLDIR` — the lint tasks in `Taskfile.yml` already do this.
