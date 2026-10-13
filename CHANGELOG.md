# Changelog

<a name="unreleased"></a>
## [0.16.1](https://github.com/eng618/go-eng/compare/v0.16.0...v0.16.1) (2026-10-13)


### Documentation

* add Copilot instructions and repository overview ([00643a5](https://github.com/eng618/go-eng/commit/00643a5acd900fc4135360d2147a9c2a0c772f4b))


### Code Refactoring

* consolidate CI workflow jobs and remove compatibility checks for web-service-gin ([8425b58](https://github.com/eng618/go-eng/commit/8425b5878f19d3ad6337014e9831c63f007e3468))
* enhance CI workflow and Taskfile to skip builds/tests for incompatible Go versions ([76f1975](https://github.com/eng618/go-eng/commit/76f1975d6553bf9a49845ad48341dfb06a0cedbf))
* **packages:** rename to avoid shadowing stdlib packages ([5d0dffb](https://github.com/eng618/go-eng/commit/5d0dffb1e9c2ebd8a2802be8a92128ccaa980ae6))
* reorganize CI workflow to separate build and test jobs, enhance linting steps, and streamline coverage upload ([56c462d](https://github.com/eng618/go-eng/commit/56c462da278b0f730a084f919726e9887388a127))
* update CI workflow to test multiple Go versions and enhance linting steps ([8a12895](https://github.com/eng618/go-eng/commit/8a128954c11b27e9a8274fa1789f6aa979bcd7e8))
* update CI workflow to use Go 1.26 and streamline linting commands ([52754a9](https://github.com/eng618/go-eng/commit/52754a90fda9909b519ec55f313437a23f82bbac))
* update codebase for modern go features, standard types, and linter standards ([144e487](https://github.com/eng618/go-eng/commit/144e487ebd339b238a6099118366122d5fa028b7))
* update Go version matrix in CI workflow to include 1.24 and 1.25 ([2fd20d3](https://github.com/eng618/go-eng/commit/2fd20d38b4b4725176a61bc41048f32323e0b88f))
* **workflow:** rename goreleaser job to release and update trigger conditions ([c033a25](https://github.com/eng618/go-eng/commit/c033a25aa6b9e1498ba8d1c823bd6626b39a4380))


### Build System

* **ci:** update Go version in CI matrix to include 1.26 ([6c575cf](https://github.com/eng618/go-eng/commit/6c575cf980055758e7dd4e779a1345ef4579ee00))
* **deps:** bump actions/checkout from 4 to 5 ([#35](https://github.com/eng618/go-eng/issues/35)) ([11f12c6](https://github.com/eng618/go-eng/commit/11f12c619b01e26c2d8862e8d9f3db373f980d8d))
* **deps:** bump actions/checkout from 5 to 6 ([45b7383](https://github.com/eng618/go-eng/commit/45b7383c3e3e1aee778387c788434cebc4991bdc))
* **deps:** bump actions/checkout from 6 to 7 ([#48](https://github.com/eng618/go-eng/issues/48)) ([a8cb9d3](https://github.com/eng618/go-eng/commit/a8cb9d3864a8f89700f3a39f6ceecd7ba0a21e10))
* **deps:** bump actions/setup-go from 5 to 6 ([#36](https://github.com/eng618/go-eng/issues/36)) ([971d737](https://github.com/eng618/go-eng/commit/971d73710c12fc19e3c202d2b21fa91f585b64b4))
* **deps:** bump actions/setup-go from 6 to 7 ([#50](https://github.com/eng618/go-eng/issues/50)) ([dbaa182](https://github.com/eng618/go-eng/commit/dbaa182303a8d0ca0b148168dc11c2c585d393db))
* **deps:** bump arduino/setup-task from 2 to 3 ([#47](https://github.com/eng618/go-eng/issues/47)) ([74d59c5](https://github.com/eng618/go-eng/commit/74d59c5659125ef5838d672b952d2c95aa2d04ab))
* **deps:** bump codecov/codecov-action from 5 to 6 ([#44](https://github.com/eng618/go-eng/issues/44)) ([36b9c97](https://github.com/eng618/go-eng/commit/36b9c97d7984ace118e7326168710e6390c2df99))
* **deps:** bump codecov/codecov-action from 6 to 7 ([#49](https://github.com/eng618/go-eng/issues/49)) ([54174fd](https://github.com/eng618/go-eng/commit/54174fd936489c58d6ed72dce6b291b88ab23cf9))
* **deps:** bump github.com/quic-go/quic-go from 0.59.0 to 0.59.1 in /examples/web-service-gin in the go_modules group across 1 directory ([#46](https://github.com/eng618/go-eng/issues/46)) ([f2d3666](https://github.com/eng618/go-eng/commit/f2d3666335f29de816678e71b72137acdc8157f1))
* **deps:** bump golangci/golangci-lint-action from 7 to 8 ([#34](https://github.com/eng618/go-eng/issues/34)) ([3715bf5](https://github.com/eng618/go-eng/commit/3715bf51188be926e11df6a710e681b1ae5818d1))
* **deps:** bump golangci/golangci-lint-action from 8 to 9 ([#38](https://github.com/eng618/go-eng/issues/38)) ([8a00ed9](https://github.com/eng618/go-eng/commit/8a00ed93ad5541e4b393433548adfbd6d5be58fc))
* **deps:** bump googleapis/release-please-action from 4 to 5 ([#45](https://github.com/eng618/go-eng/issues/45)) ([1e86d40](https://github.com/eng618/go-eng/commit/1e86d40263950e91ff459360b08dbd1b5d5e7161))
* **deps:** bump goreleaser/goreleaser-action from 6 to 7 ([ac2d443](https://github.com/eng618/go-eng/commit/ac2d443e6794d01d7e82767f9c36e50e1143ac2f))
* **deps:** downgrade Go version from 1.25 to 1.24 across all modules and configurations ([b70a57a](https://github.com/eng618/go-eng/commit/b70a57ab65208906bb54138fe5031766cc0291c5))
* **deps:** specify Go version as 1.24.0 in go.work ([99ed990](https://github.com/eng618/go-eng/commit/99ed990a40343bcd08295e862b8f7c7a2defb872))
* **deps:** update all packages ([31463c6](https://github.com/eng618/go-eng/commit/31463c6c6ca592e012aa9240651971a461664e19))
* **deps:** update Go version from 1.24 to 1.25 across modules and workflows ([32b3683](https://github.com/eng618/go-eng/commit/32b368390698af2136d680938fcc407d00fb94ec))
* **deps:** update Go version to 1.25.0 and upgrade dependencies ([53d55ae](https://github.com/eng618/go-eng/commit/53d55ae6faac9947d7fe7b7ac5821b4ff572180f))
* **deps:** update Go version to 1.26 across configuration files ([bd1f0b8](https://github.com/eng618/go-eng/commit/bd1f0b8c8958a4c173f820b501d7ceef69a093a2))


### Continuous Integration

* update dependabot config ([bb78238](https://github.com/eng618/go-eng/commit/bb78238f8e6a5163264a92a29ff5d81b290ee2c3))


### Miscellaneous Chores

* **deps:** add gopkg.in/check.v1 dependency ([4d4778b](https://github.com/eng618/go-eng/commit/4d4778b4d8d10a1eaf9ca043a4758dc2194224f4))
* migrate release and changelog generation from git-chglog to release-please ([3fac4e9](https://github.com/eng618/go-eng/commit/3fac4e978eaef84c9b3bea2fcc469eb5ad056ef2))
* PR reviewers from dependabot -&gt; CODEOWNERS ([5a3bac5](https://github.com/eng618/go-eng/commit/5a3bac597df19b03cffac9a89c1641178ae65c49))
* remove bench from validate ([c695d29](https://github.com/eng618/go-eng/commit/c695d2932c76c5171aa4020256711c33b7f7f8d8))
* replace Makefile with AGENTS.md documentation ([6a7e5c4](https://github.com/eng618/go-eng/commit/6a7e5c43d3d1a904cea9394d721102ef8a5b4670))
* update deps ([3df271c](https://github.com/eng618/go-eng/commit/3df271c369110d2c27aba4f480ea783af1b4b654))
* update Go version to 1.26 and golangci-lint to 2.11.4 across multiple modules ([3439382](https://github.com/eng618/go-eng/commit/34393824c9e56f3831477f0035862d7bc63f1086))
* update Go version to 1.26.0 and update dependencies in web-service-gin ([9228b3a](https://github.com/eng618/go-eng/commit/9228b3a8a4e63c73ae8d784fb2874d0475f79bdd))
* update Go version to 1.26.2 and bump dependencies in web-service-gin ([735baab](https://github.com/eng618/go-eng/commit/735baab72c70ec3838f0b737764436e356e6869b))
* update golangci-lint version to v2.11.4 and clarify Go version requirements in documentation ([c9b4206](https://github.com/eng618/go-eng/commit/c9b4206371220ed4b37aab4cf8bb929aba881564))
* upgrade Go to 1.27 and update toolchain versions ([c627e1d](https://github.com/eng618/go-eng/commit/c627e1dda74dca232c7c02b14ee3a9032f6191b8))

## [Unreleased]


<a name="v0.15.2"></a>
## [v0.15.2] - 2025-01-30
### Chore
- update changelog [skip ci]

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


[Unreleased]: https://github.com/eng618/go-eng/compare/v0.15.2...HEAD
[v0.15.2]: https://github.com/eng618/go-eng/compare/v0.15.1...v0.15.2
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
