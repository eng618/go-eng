# Changelog

<a name="unreleased"></a>
## [Unreleased]

### Feat
- **queue:** add benchmark tests for LinkedQueue operations and update Makefile
- **sort:** add BubbleSort algorithm with tests and documentation; remove MergeSimple function

### Fix
- **makefile:** refine benchmark output for clarity in test results


<a name="v0.15.1"></a>
## [v0.15.1] - 2025-01-30
### Feat
- **hash:** sort keys and values in hash table methods; add edge case tests and benchmarks
- **list:** implement custom linked list and wrapper for standard library; update README and tests


<a name="v0.15.0"></a>
## [v0.15.0] - 2025-01-29
### Docs
- **sort:** add package documentation for sorting algorithms and example usage

### Feat
- **fibonacci:** add Fibonacci package with multiple calculation methods and comprehensive tests
- **sort:** enhance Merge and MergeSort functions with additional edge case handling and examples

### Fix
- **goreleaser:** correct format keys in configuration and update name templates

### Test
- **binary:** add edge case tests for BinaryLoop and BinaryRecursion functions
- **linear:** add additional test cases for edge scenarios in linear search


<a name="v0.14.0"></a>
## [v0.14.0] - 2025-01-24
### Build
- update mods
- **go:** upgrade to v1.23, and apply linter fixes

### Docs
- improve documentation for main package and Big O notation; enhance comments in SliceQueue implementation
- update README and add documentation for anagrams and hash packages; enhance list, queue, and stack documentation
- **SECURITY:** enhance vulnerability reporting guidelines and provide detailed instructions
- **anagrams:** improve function documentation and add example usage; enhance test coverage
- **stack:** enhance documentation for New, NewSeeded, and Pop methods; add concurrency tests

### Fix
- **stack:** remove redundant lock in Pop method; update benchmarks to ignore error checks

### Refactor
- **circular:** remove unused circular package and associated tests
- **linkedList:** update node struct to export fields and improve documentation

### Test
- **hash:** add test coverage and docs
- **linkedList:** fix race condition


<a name="v0.13.4"></a>
## [v0.13.4] - 2025-01-06
### Build
- **deps:** bump all
- **deps:** bump codecov/codecov-action from 4 to 5 ([#28](https://github.com/eng618/go-eng/issues/28))

### Fix
- **lint:** correct redefinition

### Test
- update tests, and bump go version


<a name="v0.13.3"></a>
## [v0.13.3] - 2024-06-13
### Ci
- **actions:** update gorelease


<a name="v0.13.2"></a>
## [v0.13.2] - 2024-06-13
### Build
- **deps:** bump golangci/golangci-lint-action from 5 to 6 ([#27](https://github.com/eng618/go-eng/issues/27))

### Chore
- update changelog [skip ci]


<a name="v0.13.1"></a>
## [v0.13.1] - 2024-06-13
### Ci
- **goreleaser:** update version


<a name="v0.13.0"></a>
## [v0.13.0] - 2024-06-13
### Build
- bump tools and deps
- **go:** upgrade to v1.22

### Ci
- only 1.22
- check latest go version

### Reverts
- ci: only 1.22
- ci: check latest go version


<a name="v0.12.2"></a>
## [v0.12.2] - 2024-05-02
### Build
- bump deps
- **deps:** bump golangci/golangci-lint-action from 3 to 5 ([#26](https://github.com/eng618/go-eng/issues/26))
- **deps:** bump codecov/codecov-action from 3 to 4 ([#22](https://github.com/eng618/go-eng/issues/22))
- **deps:** bump golang.org/x/crypto in /examples/web-service-gin
- **deps:** bump actions/setup-go from 4 to 5 ([#21](https://github.com/eng618/go-eng/issues/21))

### Chore
- update changelog [skip ci]

### Test
- **db:** updates
- **db:** lint tests


<a name="v0.12.1"></a>
## [v0.12.1] - 2023-12-19
### Build
- **makefile:** update depricated flag

### Chore
- update changelog [skip ci]

### Feat
- **db:** return set time

### Test
- **db:** add test coverage


<a name="v0.12.0"></a>
## [v0.12.0] - 2023-12-18
### Build
- remove disabled workflow
- **deps:** bump all deps
- **deps:** bump all

### Chore
- update changelog [skip ci]

### Docs
- **grammar:** updates phonetics

### Feat
- **algo:** adds anagrams package
- **db:** add in memory database example
- **leet:** add a couple answers
- **list:** standard library example
- **vowels:** add vowels algorithm package

### Fix
- **anagrams:** simplified logic

### Refactor
- **fibonacci:** clean up typos and examples
- **list:** standard library example

### Test
- remove expected output
- **slice_queue:** update examples


<a name="v0.11.8"></a>
## [v0.11.8] - 2023-10-02
### Build
- **deps:** bump actions/checkout from 3 to 4 ([#18](https://github.com/eng618/go-eng/issues/18))
- **deps:** bump goreleaser/goreleaser-action from 4 to 5 ([#17](https://github.com/eng618/go-eng/issues/17))

### Ci
- disable twitter announce


<a name="v0.11.7"></a>
## [v0.11.7] - 2023-08-18
### Build
- go1.21

### Ci
- fix action version


<a name="v0.11.6"></a>
## [v0.11.6] - 2023-08-18
### Ci
- update releaser config
- update branches


<a name="v0.11.5"></a>
## [v0.11.5] - 2023-08-13

<a name="v0.11.4"></a>
## [v0.11.4] - 2023-08-13
### Build
- **mod:** update mod


<a name="v0.11.3"></a>
## [v0.11.3] - 2023-06-10
### Chore
- update changelog [skip ci]

### Docs
- remove unneeded README

### Feat
- **examples:** add RESTFul API using Gin

### Test
- add simple test


<a name="v0.11.2"></a>
## [v0.11.2] - 2023-06-10
### Fix
- correct generics path name

### Revert
- change package name back to github hosted path ([#16](https://github.com/eng618/go-eng/issues/16))


<a name="v0.11.1"></a>
## [v0.11.1] - 2023-06-07
### Build
- bump go version

### Chore
- minor updates
- update changelog [skip ci]

### Docs
- add package documentation

### Feat
- add generics example ([#14](https://github.com/eng618/go-eng/issues/14))


<a name="v0.11.0"></a>
## [v0.11.0] - 2023-05-06
### Chore
- update changelog [skip ci]

### Ci
- update release workflow

### Fix
- fully update module name


<a name="v0.10.0"></a>
## [v0.10.0] - 2023-05-06
### Chore
- add publish command

### Ci
- remove support for go 1.18

### Feat
- update module name

### Fix
- cleanup go.mod


<a name="v0.9.2"></a>
## [v0.9.2] - 2023-04-07
### Ci
- update goreleaser


<a name="v0.9.1"></a>
## [v0.9.1] - 2023-04-07
### Chore
- minor adjustments


<a name="v0.9.0"></a>
## [v0.9.0] - 2023-04-07
### Build
- upgrade go to 1.19 latest
- bump go patch version
- **deps:** bump actions/setup-go from 3 to 4

### Chore
- update changelog [skip ci]
- update changelog [skip ci]
- update changelog template
- update changelog [skip ci]
- update changelog template
- update changelog [skip ci]
- update changelog comand
- update changelog [skip-ci]
- update changelog comand
- update changelog
- update changelog comand

### Ci
- comment out deprications
- update goreleaser
- fix go version
- add dependabot.yml
- update lint tool
- udpate to go1.20
- remove go 1.20
- add go 1.20 and update releaser version
- update actions

### Docs
- update documentation

### Feat
- go 1.20

### Test
- speed up tests un nanosecond vs second

### Pull Requests
- Merge pull request [#13](https://github.com/eng618/go-eng/issues/13) from eng618/dependabot/github_actions/actions/setup-go-4


<a name="v0.8.1"></a>
## [v0.8.1] - 2022-09-11
### Build
- add changelog command

### Chore
- update changelog
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


[Unreleased]: https://github.com/eng618/go-eng/compare/v0.15.1...HEAD
[v0.15.1]: https://github.com/eng618/go-eng/compare/v0.15.0...v0.15.1
[v0.15.0]: https://github.com/eng618/go-eng/compare/v0.14.0...v0.15.0
[v0.14.0]: https://github.com/eng618/go-eng/compare/v0.13.4...v0.14.0
[v0.13.4]: https://github.com/eng618/go-eng/compare/v0.13.3...v0.13.4
[v0.13.3]: https://github.com/eng618/go-eng/compare/v0.13.2...v0.13.3
[v0.13.2]: https://github.com/eng618/go-eng/compare/v0.13.1...v0.13.2
[v0.13.1]: https://github.com/eng618/go-eng/compare/v0.13.0...v0.13.1
[v0.13.0]: https://github.com/eng618/go-eng/compare/v0.12.2...v0.13.0
[v0.12.2]: https://github.com/eng618/go-eng/compare/v0.12.1...v0.12.2
[v0.12.1]: https://github.com/eng618/go-eng/compare/v0.12.0...v0.12.1
[v0.12.0]: https://github.com/eng618/go-eng/compare/v0.11.8...v0.12.0
[v0.11.8]: https://github.com/eng618/go-eng/compare/v0.11.7...v0.11.8
[v0.11.7]: https://github.com/eng618/go-eng/compare/v0.11.6...v0.11.7
[v0.11.6]: https://github.com/eng618/go-eng/compare/v0.11.5...v0.11.6
[v0.11.5]: https://github.com/eng618/go-eng/compare/v0.11.4...v0.11.5
[v0.11.4]: https://github.com/eng618/go-eng/compare/v0.11.3...v0.11.4
[v0.11.3]: https://github.com/eng618/go-eng/compare/v0.11.2...v0.11.3
[v0.11.2]: https://github.com/eng618/go-eng/compare/v0.11.1...v0.11.2
[v0.11.1]: https://github.com/eng618/go-eng/compare/v0.11.0...v0.11.1
[v0.11.0]: https://github.com/eng618/go-eng/compare/v0.10.0...v0.11.0
[v0.10.0]: https://github.com/eng618/go-eng/compare/v0.9.2...v0.10.0
[v0.9.2]: https://github.com/eng618/go-eng/compare/v0.9.1...v0.9.2
[v0.9.1]: https://github.com/eng618/go-eng/compare/v0.9.0...v0.9.1
[v0.9.0]: https://github.com/eng618/go-eng/compare/v0.8.1...v0.9.0
[v0.8.1]: https://github.com/eng618/go-eng/compare/v0.8.0...v0.8.1
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
