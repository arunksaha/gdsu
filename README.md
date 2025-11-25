# gdsu — Generic Disjoint Set Union (Union-Find)

`gdsu` is a modern, ergonomic, and extensible **Disjoint Set Union (DSU)** / **Union-Find** library for Go.

It provides:

- A clean **interface-based abstraction** for DSU operations  
- A **sparse**, **generic**, **map-backed** implementation (unbounded capacity)  
  - Support for any `comparable` type
- A **compact**, integer-indexed, slice-backed implementation (fixed capacity)  
  - High-performance compact mode for integer-indexed workloads  

This makes `gdsu` suitable for:
- Graph algorithms  
- Connectivity problems  
- Clustering  
- Dynamic merging/partitioning tasks  
- Use with arbitrary user-defined types  

---

# 1. DSU Interface

Located in: `gdsu/gdsu.go`

The interface defines the core DSU behavior:

```go
type DSU[T comparable] interface {
    Find(x T) T
    Union(x, y T) bool
    Connected(x, y T) bool
    Groups() map[T][]T
}
```

Both implementations, sparse and compact, satisfy this interface.

---

# 2. Sparse DSU (Generic, Map-Based, Unbounded)

Located in: `sparse/sparse.go`

### ✨ Features
- Generic over any `comparable` type  
- Grows dynamically — **no fixed initial size**  
- Backed by Go maps (`map[T]T`)  
- Ideal for:
  - Dynamic workloads  
  - Arbitrary user-defined element types  
  - Situations where IDs are not known upfront  

### Example

```go
import "github.com/arunksaha/gdsu/sparse"

dsu := sparse.New[string]()

dsu.Union("apple", "banana")
dsu.Union("carrot", "banana")

fmt.Println(dsu.Connected("apple", "carrot")) // true
fmt.Println(dsu.Groups())
// map[banana:[apple banana carrot]]
```

---

# 3. Compact DSU (Int-Indexed, Slice-Backed)

Located in: `compact/compact.go`

### ✨ Features
- High-performance, low-overhead  
- Slice-based parent and rank arrays  
- Requires a fixed capacity specified at initialization  
- Ideal for:
  - Graph algorithms with node IDs `0..N-1`  
  - Tight loops  
  - Performance-critical workloads  

### Example

```go
import "github.com/arunksaha/gdsu/compact"

dsu := compact.New(10) // supports elements 0..9

dsu.Union(1, 2)
dsu.Union(2, 3)

fmt.Println(dsu.Connected(1, 3)) // true
fmt.Println(dsu.Groups())
// map[1:[1 2 3]]
```

---

# 4. Unique offerings of gdsu

Often Union-Find packages
- Support **only integers**  
- Require **fixed capacity** at initialization  
- Do not support **generic types**  
- Do not provide an interface abstraction  

`gdsu` provides:

### ✔ Generic sparse DSU  
Supports **any comparable type** and grows dynamically.

### ✔ Unified interface  
Both implementations satisfy the same `DSU[T]` interface.

### ✔ Two optimized backends  
- **Sparse**: flexible, dynamic, generic  
- **Compact**: fast, contiguous memory, ideal for high performance workloads  

### ✔ Clean package layout  
- `gdsu/` — interface  
- `sparse/` — generic dynamic implementation  
- `compact/` — integer-slice fixed-capacity implementation  

---

# 5. Project Structure

```
gdsu/
  gdsu.go
sparse/
  sparse.go
  sparse_test.go
compact/
  compact.go
  compact_test.go
```

---

# 6. License

MIT License — see `LICENSE`.

---

# 7. Contributions

PRs and issues welcome!
