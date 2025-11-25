package compact

import "github.com/arunksaha/gdsu"

// DSU is an int-based, slice-backed Disjoint-Set Union implementation.
//
// It manages integers in the fixed range [0, n). It is compact and fast
// but not sparse: you must choose the capacity at construction time.
type DSU struct {
	parent []int
	rank   []int
}

// New creates a DSU for elements in the range [0, size).
func New(size int) *DSU {
	if size < 0 {
		size = 0
	}
	parent := make([]int, size)
	rank := make([]int, size)
	for i := 0; i < size; i++ {
		parent[i] = i
		rank[i] = 0
	}
	return &DSU{
		parent: parent,
		rank:   rank,
	}
}

// boundsCheck ensures x is within [0, len(parent)).
func (dsu *DSU) boundsCheck(x int) bool {
	return x >= 0 && x < len(dsu.parent)
}

// Find returns the representative element (root) of the set containing x.
// If x is out of range, it panics (by design for this compact DSU).
func (dsu *DSU) Find(x int) int {
	if !dsu.boundsCheck(x) {
		panic("compact.DSU: index out of range in Find")
	}

	root := x
	// first walk to root
	for dsu.parent[root] != root {
		root = dsu.parent[root]
	}

	// compress
	for x != root {
		p := dsu.parent[x]
		dsu.parent[x] = root
		x = p
	}

	return root
}

// Union merges the sets containing x and y.
// Returns true if the sets were separate and are now merged.
// Panics if x or y are out of range.
func (dsu *DSU) Union(x, y int) bool {
	if !dsu.boundsCheck(x) || !dsu.boundsCheck(y) {
		panic("compact.DSU: index out of range in Union")
	}
	rootX, rootY := dsu.Find(x), dsu.Find(y)
	if rootX == rootY {
		return false
	}
	if dsu.rank[rootX] < dsu.rank[rootY] {
		rootX, rootY = rootY, rootX
	}
	dsu.parent[rootY] = rootX
	if dsu.rank[rootX] == dsu.rank[rootY] {
		dsu.rank[rootX]++
	}
	return true
}

// Connected reports whether x and y are in the same set.
// Panics if x or y are out of range.
func (dsu *DSU) Connected(x, y int) bool {
	if !dsu.boundsCheck(x) || !dsu.boundsCheck(y) {
		panic("compact.DSU: index out of range in Connected")
	}
	return dsu.Find(x) == dsu.Find(y)
}

// Groups returns a map from root -> slice of elements in that set.
func (dsu *DSU) Groups() map[int][]int {
	groups := make(map[int][]int)
	for x := range dsu.parent {
		root := dsu.Find(x)
		groups[root] = append(groups[root], x)
	}
	return groups
}

// Compile-time assertion that DSU implements gdsu.DSU[int].
var _ gdsu.DSU[int] = (*DSU)(nil)
