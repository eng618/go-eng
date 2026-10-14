# Cut a Release

How releases work. Releases are automated — your job is the commit message.

## 1. Use conventional commits

Every commit message starts with a type: `feat:`, `fix:`, `chore:`, and so on. The type decides the version bump:

| Commit                          | Bump  |
| ------------------------------- | ----- |
| `feat:`                         | minor |
| `fix:`                          | patch |
| any other type                  | patch |
| `!` or `BREAKING CHANGE` footer | major |

Example:

```sh
git commit -m "feat: add insertion sort"
```

## 2. Push and let release-please work

`release-please` owns `CHANGELOG.md` — never edit it by hand. Merging to the default branch with correctly-typed commits produces the release PR and, on merge, the tagged release.

## 3. Before pushing

```sh
task validate
```

See [Run tasks](run-tasks.md). CI (`ci.yml`) runs `task lint`, then `task test`, then `task build` with a 30-minute timeout.
