[![Go Reference](https://pkg.go.dev/badge/github.com/arunksaha/gdsu.svg)](https://pkg.go.dev/github.com/arunksaha/gdsu)
[![Go Report Card](https://goreportcard.com/badge/github.com/arunksaha/gdsu)](https://goreportcard.com/report/github.com/arunksaha/gdsu)
![Build](https://github.com/arunksaha/gdsu/actions/workflows/ci.yml/badge.svg)
[![codecov](https://codecov.io/github/arunksaha/gdsu/graph/badge.svg?token=134TP2MY41)](https://codecov.io/github/arunksaha/gdsu)


# gdsu — Generic Disjoint Set Union (Union-Find)

`gdsu` is a modern, type-safe, composable Disjoint Set Union (DSU) / Union-Find Go library.  
It provides a clean DSU interface with two interchangeable implementations:

- **Sparse DSU** — generic, map-backed, supports any comparable type, grows dynamically  
- **Compact DSU** — high-performance integer-indexed version using slices, fixed capacity  

This library is designed for algorithmic workloads, data pipelines, graph theory, clustering, and any use-case involving connectivity queries.

---

## 1. DSU Interface

Located in: `gdsu.go`

```go
type DSU[T comparable] interface {
    Find(x T) T            // Find returns the representative element (root) of x
    Union(x, y T) bool     // Union merges the sets containing x and y
    Connected(x, y T) bool // Connected reports whether x and y are in the same set
    Groups() map[T][]T     // Groups returns all connected components
}
```

Both implementations — sparse and compact — satisfy this interface.

---

## 2. Sparse DSU (Generic, Map-Based, Dynamic)

Located in: `sparse/sparse.go`

### Key Features
- Works with **any comparable type**: strings, integers, structs, user-defined keys  
- Supports dynamic growth — **no fixed initial capacity**
- Backed by Go maps (`map[T]T`)
- Ideal for:
  - Arbitrary keys  
  - Sparse connectivity  
  - Unpredictable element ranges  
  - Rapid prototyping with rich data types  

### Example

```go
package main

import (
    "fmt"
    "github.com/arunksaha/gdsu/sparse"
)

func main() {
    dsu := sparse.New[string]()

    dsu.Union("apple", "banana")
    dsu.Union("banana", "cherry")

    fmt.Println(dsu.Connected("apple", "cherry")) // true
    fmt.Println(len(dsu.Groups()))                // 1
}
```

---

## 3. Compact DSU (Int-Indexed, Slice-Backed)

Located in: `compact/compact.go`

### Key Features
- Very fast, minimal overhead  
- Uses contiguous slices for parent and rank  
- Requires fixed capacity at initialization (`New(size)`)
- Ideal for:
  - Graph algorithms  
  - Tight inner loops  
  - Integer node IDs (`0..N-1`)  
  - Performance-critical workloads  

### Example

```go
package main

import (
    "fmt"
    "github.com/arunksaha/gdsu/compact"
)

func main() {
    dsu := compact.New(10) // supports 0..9

    dsu.Union(1, 2)
    dsu.Union(2, 3)

    fmt.Println(dsu.Connected(1, 3)) // true
    fmt.Println(dsu.Groups())
}
```

---

## 4. Benchmark Overview

Benchmark files:
- `sparse/sparse_benchmark_test.go`
- `compact/compact_benchmark_test.go`
- `comparison/comparison_benchmark_test.go`

Benchmarks include:
- Union performance  
- Find performance  
- Connected queries  
- Mixed-operation simulations  
- Sparse vs compact comparison  
- Memory profiling  

Compact is optimized for pure speed; sparse is optimized for flexibility.

---

## 5. Package Structure

```
.
├── compact
│   ├── compact_benchmark_test.go
│   ├── compact_example_test.go
│   ├── compact.go
│   └── compact_test.go
├── comparison
│   └── comparison_benchmark_test.go
├── gdsu.go
├── gdsu_test.go
├── go.mod
├── LICENSE
├── Makefile
├── README.md
└── sparse
    ├── sparse_benchmark_test.go
    ├── sparse_example_test.go
    ├── sparse.go
    └── sparse_test.go
```

---

## 6. Unique aspects of gdsu

Often (Go) DSU implementations:

- Work only with `int`
- Require fixed capacity specified at construction
- Do not support generic types 
- Do not provide an interface abstraction

`gdsu` improves on all of these:

### ✔ Unified interface  
Both implementations satisfy the same `DSU[T]` interface.

### ✓ Generic sparse DSU  
Supports any comparable type, grows dynamically.

### ✓ High-performance compact DSU  
Optimized for numeric workloads.

---

## 7. Installation

```sh
go get github.com/arunksaha/gdsu
```

---

## 8. Getting Started

Choose sparse or compact depending on your needs.  
Example usage is provided in each subpackage’s `*_example_test.go`.

The `Makefile` provides several convenience targets, namely:

  - `lint` to run various formatting and static analysis tools, e.g., `fmt`, `vet`, `staticcheck`, `tidy`

  - `test` to run unit tests

  - `coverage` to test coverage

  - `covhtml` to view test coverage results in pretty color-coded html

  - `doc` to generate documentation for local viewing

  - `benchmark` to run all the benchmark tests

---

## 9. Contributions

Issues and PRs are welcome!
