# Modules

The workspace (`go.work`) and its modules. `go` commands for a submodule must run in that directory; the cross-module equivalent from the root is `task <module>-<action>`.

| Module          | Directory                         | Import path / notes                                                           |
| --------------- | --------------------------------- | ----------------------------------------------------------------------------- |
| Root            | `.`                               | `github.com/eng618/go-eng` — `algo/`, `ds/`, `lib/`; `main.go` is a demo only |
| Circuit breaker | `design-patterns/circuit-breaker` | Separate module                                                               |
| Master-slave    | `design-patterns/master-slave`    | Separate module                                                               |
| Contexts        | `examples/contexts`               | Separate module                                                               |
| Generics        | `examples/generics`               | Separate module                                                               |
| Web service     | `examples/web-service-gin`        | Separate module                                                               |

Go version: 1.27 (`.tool-versions` is the source of truth).
