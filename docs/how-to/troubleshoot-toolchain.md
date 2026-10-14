# Troubleshoot the Toolchain

How to fix the known environment quirks in this repo.

## Symptom: `compile: version "go1.26.2" does not match go tool version "go1.27.2"`

Your shell exports a stale `GOROOT` (e.g. from `asdf`) pointing at an older Go installation than the active toolchain. Confirm:

```sh
echo "GOROOT=$GOROOT"
```

If it points at a version other than the one in `.tool-versions` (Go 1.27), unset it for the failing command:

```sh
env -u GOROOT -u GOTOOLDIR go test ./algo/sorting/...
```

`Taskfile.yml` already wraps all lint tasks this way (`env -u GOROOT -u GOTOOLDIR golangci-lint run ./...`), so `task lint` is unaffected. Only raw `go` / `golangci-lint` invocations need the prefix. To fix your shell permanently, update or remove the `GOROOT` export in your profile.

## Rule: never run bare `golangci-lint run ./...`

Always go through `task lint` (or copy its `env -u GOROOT -u GOTOOLDIR` prefix). The bare invocation hits the quirk above.

## `GOTOOLCHAIN=local`

`Taskfile.yml` sets `GOTOOLCHAIN=local`, pinning builds to the installed toolchain instead of auto-downloading. Keep it — and keep `.tool-versions` as the source of truth for `golang` / `golangci-lint` versions (update it together with `.github/workflows/ci.yml`).

## Tests need CGO for `-race`

The standard test invocation is:

```sh
CGO_ENABLED=1 go test ./... -cover -race -coverprofile=coverage.out -covermode=atomic
```

`task test` runs this per module. `coverage.out` and `linter-output.json` are generated and gitignored — do not commit them.
