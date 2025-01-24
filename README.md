# go-eng

[![build-test-lint](https://github.com/eng618/go-eng/actions/workflows/ci.yml/badge.svg)](https://github.com/eng618/go-eng/actions/workflows/ci.yml)
[![release](https://github.com/eng618/go-eng/actions/workflows/release.yml/badge.svg)](https://github.com/eng618/go-eng/actions/workflows/release.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/eng618/go-eng)](https://goreportcard.com/report/github.com/eng618/go-eng)
[![codecov](https://codecov.io/gh/eng618/go-eng/branch/main/graph/badge.svg?token=L1D7JSRK98)](https://codecov.io/gh/eng618/go-eng)

## Overview

This is a personal package I will be using to create various helper packages,
as I learn and explore the Go programming language. The main package of this
repo is simply to demonstrate the use of a subset of the contained packages.

For full documentation visit the individual packages listed in the directories
section [pkg.go.dev/github.com/eng618/go-eng#section-directories](https://pkg.go.dev/github.com/eng618/go-eng#section-directories).

## List of Packages

### Algorithms

- `github.com/eng618/go-eng/algo/anagrams`
  - anagram detection
- `github.com/eng618/go-eng/algo/search`
  - binary search
    - loop implementation
    - recursion implementation
  - linear search
- `github.com/eng618/go-eng/algo/sort`
  - merge sort
  - insertion sort
  - bubble sort
  - quick sort

### Data Structures

- `github.com/eng618/go-eng/ds/list`
  - linked list
  - double linked list
- `github.com/eng618/go-eng/ds/queue`
  - slice implementation
  - linked list implementation
- `github.com/eng618/go-eng/ds/stack`
  - slice implementation
- `github.com/eng618/go-eng/ds/hash`
- `github.com/eng618/go-eng/algo/circular`
  - circular list

## Still to come

### Planned Algorithms

- Recursion
- Sort

  - [ ] Insertion
  - [ ] Bubble
  - [ ] Quick sort

- Search

  - [ ] Breadth-first search (BFS)
  - [ ] Depth-first search (DFS)

### Planned Data Structures

- [x] Queue
- [ ] Double linked list
- [ ] Binary Search Tree (BST)
- [ ] Graph
