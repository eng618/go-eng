# Changelog

<a name="unreleased"></a>

## [Unreleased]

### Docs

- updated README, add module doc

### Feat

- **queue:** add LinkedList implementation

<a name="v0.4.3"></a>

## [v0.4.3] - 2021-06-14

### Feat

- **linkedlist:** add Remove method

<a name="v0.4.2"></a>

## [v0.4.2] - 2021-06-14

### Feat

- enhanced merge sort
  - made search a single package
  - **queue:** add slice based queue Example
  - **search:** add linear function

### Refactor

- more changes for readability
  - remove Data, and simply define []int
  - **linkedlist:** to be more concise

### Test

- **benchmark:** add benchmark tests

<a name="v0.4.1"></a>

## [v0.4.1] - 2021-06-11

### Ci

- fix creds

<a name="v0.4.0"></a>

## [v0.4.0] - 2021-06-11

### Ci

- add twitter creds

<a name="v0.3.0"></a>

## [v0.3.0] - 2021-06-11

### Ci

- only run goreleaser on tags

<a name="v0.2.0"></a>

## [v0.2.0] - 2021-06-11

### Refactor

- organized for ease of use

<a name="v0.1.0"></a>

## v0.1.0 - 2021-06-11

### Build

- add goreleaser and required configuration.

### Ci

- add actions to build, test, and lint

### Docs

- add Big O cheatsheet
  - update readme

### Feat

- add merge sort package
  - add main
  - **binary:** add binary search package
  - **linkedlist:** create linkedlist package
  - **stack:** add stack package

### Fix

- various fixes...
  - correct go mod name, and add go reportcard badge

### Refactor

- replaced how new stacks are created
  - use go naming conventions for package

### Test

- add tests for merge sort
  - add test cases for list
  - fix delete tests
  - add delete tests
  - fix example output

[unreleased]: https://github.com/eng618/go-eng/compare/v0.4.3...HEAD
[v0.4.3]: https://github.com/eng618/go-eng/compare/v0.4.2...v0.4.3
[v0.4.2]: https://github.com/eng618/go-eng/compare/v0.4.1...v0.4.2
[v0.4.1]: https://github.com/eng618/go-eng/compare/v0.4.0...v0.4.1
[v0.4.0]: https://github.com/eng618/go-eng/compare/v0.3.0...v0.4.0
[v0.3.0]: https://github.com/eng618/go-eng/compare/v0.2.0...v0.3.0
[v0.2.0]: https://github.com/eng618/go-eng/compare/v0.1.0...v0.2.0
