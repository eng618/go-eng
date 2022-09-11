# Changelog

<a name="unreleased"></a>

## [Unreleased]

### Build

- add changelog command
  
  ### Chore

- update change log
  
  
<a name="v0.8.0"></a>

## [v0.8.0] - 2022-09-11

### Build

- upgrade go to 1.18
  - update to latest go 1.17.x
  - update release command
  
  ### Feat

- add package to write to a file
  
  ### Style

- apply formatting
  
  
<a name="v0.7.0"></a>

## [v0.7.0] - 2022-05-06

### Build

- add release command
  - update some Makefile commands
  - update remaining 1.16 references
  - bump default build version to 1.17
  
  ### Chore

- create Makefile
  
  ### Ci

- update lint configuration
  - remove verbose test logging
  - test go 1.18 and 1.17 only
  - specify coverage file
  - update codecov token
  - update golangci config
  - only use go version n-2
  - add code coverage with Codecov
  
  ### Docs

- add codecov badge
  - update CHANGELOG
  
  ### Feat

- add context with timeout examples
  - stub circular package
  
  
<a name="v0.6.0"></a>

## [v0.6.0] - 2021-06-17

### Docs

- update changelog [skip-ci]
  
  ### Feat

- **fibonacci:** add algorithms to calculate fib
  - **queue:** add Peek method, increase test cov
  
  
<a name="v0.5.0"></a>

## [v0.5.0] - 2021-06-16

### Chore

- **lint:** correct typos and golint warnings
  
  ### Docs

- add CHANGELOG
  - updated README, add module doc
  
  ### Feat

- **hash:** add hash table data structure
  - **queue:** add LinkedList implementation
  
  ### Refactor

- integrated golangci-lint with config
  
  
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
  
  
[Unreleased]: https://github.com/eng618/go-eng/compare/v0.8.0...HEAD
[v0.8.0]: https://github.com/eng618/go-eng/compare/v0.7.0...v0.8.0
[v0.7.0]: https://github.com/eng618/go-eng/compare/v0.6.0...v0.7.0
[v0.6.0]: https://github.com/eng618/go-eng/compare/v0.5.0...v0.6.0
[v0.5.0]: https://github.com/eng618/go-eng/compare/v0.4.3...v0.5.0
[v0.4.3]: https://github.com/eng618/go-eng/compare/v0.4.2...v0.4.3
[v0.4.2]: https://github.com/eng618/go-eng/compare/v0.4.1...v0.4.2
[v0.4.1]: https://github.com/eng618/go-eng/compare/v0.4.0...v0.4.1
[v0.4.0]: https://github.com/eng618/go-eng/compare/v0.3.0...v0.4.0
[v0.3.0]: https://github.com/eng618/go-eng/compare/v0.2.0...v0.3.0
[v0.2.0]: https://github.com/eng618/go-eng/compare/v0.1.0...v0.2.0
