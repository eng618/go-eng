# AGENTS.md

Learning/example Go repo — every public API is taught, not just shipped: explain it in `doc.go` and demonstrate it with the BET method (Benchmark + Example + Test) in `*_test.go`.

Go workspace (`go.work`) — root module plus 6 independent modules. `Taskfile.yml` is the source of truth for all workflows.

## Modules

| Dir | Notes |
| --- | --- |
| `.` (root `github.com/eng618/go-eng`) | `algo/`, `ds/`, `lib/`; `main.go` is a demo only, not a library entrypoint |
| `design-patterns/circuit-breaker` | Separate module |
| `design-patterns/master-slave` | Separate module |
| `examples/contexts`, `examples/generics`, `examples/rate-limiting`, `examples/web-service-gin` | Separate modules each |

`go` commands for a submodule must run in that dir. Cross-module equivalent is `task <module>:<action>` from the root (e.g. `task circuit-breaker:test`, `task web-service-gin:build`), or `task <action>` from inside the module dir. Each module owns its `Taskfile.yml`, included by the root `Taskfile.yml` (root module tasks live in `Taskfile.root.yml`).

## Commands

```sh
task validate   # required order: lint -> test -> build (mirrors CI)
task lint       # all modules; fix with task lint-fix
task test       # all modules: CGO_ENABLED=1 go test ./... -cover -race -coverprofile=coverage.out -covermode=atomic
task build      # all modules
task root:test   # root module only
go test ./algo/sorting/... -run TestBubbleSort            # single test (from module dir)
task bench      # benchmarks, all modules
go test ./algo/sorting/... -bench=BenchmarkBubbleSort     # single benchmark
```

Do not run bare `golangci-lint run ./...` or bare `go` commands; Taskfile wraps them as `env -u GOROOT -u GOTOOLDIR ...` (toolchain quirk: the asdf golang plugin exports a `GOROOT` that can mismatch the shim-resolved `go` binary, failing the build with `compile: version ... does not match go tool version ...`). `Taskfile` sets `GOTOOLCHAIN=local`.

## Toolchain

- Go 1.27. `.tool-versions` (`asdf`/`mise`) is source of truth for `golang` / `golangci-lint` versions — update it and `.github/workflows/ci.yml` together.
- CI (`ci.yml`): `task lint`, then `task test`, then `task build`. Timeout 30m, concurrency cancels in-progress runs on same ref.

## Conventions (enforced by `.golangci.yml`)

- Imports (`gci`): stdlib, blank line, third-party, blank line, `github.com/eng618/go-eng`. `custom-order: true`. Module paths must stay dotted — gofumpt sorts dotless imports as stdlib and fights `gci` grouping.
- Formatters: `gofumpt` (with `extra` group-params/clothe-returns/balance-calls), `gci`, `goimports`. No `gofmt` (subset of gofumpt), no `golines` — line length is `lll` at 120, report-only.
- Extra linters (teaching set): `modernize`, `testableexamples`, `testpackage`, `thelper`, `godoclint`, `intrange`, `perfsprint`, `usestdlibvars`. White-box tests live in `*_internal_test.go` (the only `testpackage` exemption); all other tests are external `foo_test`.
- Tests (BET, required for every public API): external package (`package foo_test`), table-driven with `t.Run(tt.name, ...)`; an `Example*` func that teaches usage (verified by `go test` via `// Output:`); a `Benchmark*` func even when perf is not the point.
- Every package has a `doc.go` holding the top-level package documentation; new packages must add one (follow the existing pattern).
- Conventional commits (`feat:`, `fix:`, `chore:`, ...): `feat` → minor, `fix` → patch, any other type → patch, `!`/`BREAKING CHANGE` → major. `CHANGELOG.md` is owned by release-please — never edit it by hand.

## Artifacts

- `linter-output.json` and `coverage.out` are generated and gitignored — do not commit.
