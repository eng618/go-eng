# Documentation

Start here. These docs follow the [Diátaxis](https://diataxis.fr) framework: four kinds of documentation for four kinds of need.

| Need                             | Go to                                     | Contains                                                      |
| -------------------------------- | ----------------------------------------- | ------------------------------------------------------------- |
| I want to **learn** by doing     | [Tutorials](tutorials/getting-started.md) | Step-by-step labs that assume no prior knowledge of this repo |
| I want to **get something done** | [How-to guides](how-to/run-tasks.md)      | Task-focused recipes: run, test, add code, release            |
| I want to **look something up**  | [Reference](reference/packages.md)        | Dry, complete indexes: modules, packages, commands            |
| I want to **understand**         | [Explanation](explanation/big-o.md)       | Concepts and design decisions behind the code                 |

Project status (implemented vs. planned) lives alongside, outside the four quadrants: [Implementation Inventory](project/inventory.md).

## Tutorials

1. [Getting Started](tutorials/getting-started.md) — clone, run the demo, run your first test.
2. [Sorting Lab](tutorials/sorting-lab.md) — sort data hands-on with `algo/sorting`.
3. [Data Structures Lab](tutorials/data-structures-lab.md) — stack, queue, and linked list in action.

## How-to guides

- [Run tasks](how-to/run-tasks.md) — `task validate`, `lint`, `test`, `build`, `bench`, per-module runs.
- [Add a package](how-to/add-a-package.md) — `doc.go` + BET checklist for new algorithms or data structures.
- [Troubleshoot the toolchain](how-to/troubleshoot-toolchain.md) — GOROOT quirk, `GOTOOLCHAIN=local`, CGO test flags.
- [Cut a release](how-to/cut-a-release.md) — conventional commits and release-please.

## Reference

- [Modules](reference/modules.md) — workspace module table.
- [Packages](reference/packages.md) — package index with links to `doc.go` and pkg.go.dev.
- [Commands](reference/commands.md) — Taskfile command lookup.

## Explanation

- [Big-O cheatsheet](explanation/big-o.md) — complexity rules used across the labs.
- [Design decisions](explanation/design-decisions.md) — why a workspace, why BET, why Taskfile.
