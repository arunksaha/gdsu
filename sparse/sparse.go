package sparse

import "github.com/arunksaha/gdsu"

// DSU is a sparse, map-backed generic implementation of a Disjoint-Set Union.
//
// It does NOT require a fixed capacity or pre-registration of elements.
// Elements are added lazily when first seen by Find/Union.
type DSU[T comparable] struct {
	// parent stores the immediate parent of each element;
	// if parent[x] == x, then x is the root of its set.
	parent map[T]T

	// rank stores an upper bound on the height of the tree rooted at each element.
	// Used with union-by-rank to keep trees shallow and operations near O(1).
	rank map[T]int
}

// New creates a new DSU initialized with the given elements.
// Additional elements may still be added later via Find/Union.
func New[T comparable](elems ...T) *DSU[T] {
	dsu := &DSU[T]{
		parent: make(map[T]T, len(elems)),
		rank:   make(map[T]int, len(elems)),
	}
	for _, e := range elems {
		dsu.parent[e] = e
		dsu.rank[e] = 0
	}
	return dsu
}

// Find returns the representative element (root) of the set containing x.
// If x is not present, it is added as a singleton set.
func (dsu *DSU[T]) Find(x T) T {
	// if unseen, initialize
	if _, ok := dsu.parent[x]; !ok {
		dsu.parent[x] = x
		dsu.rank[x] = 0
		return x
	}

	// find root
	root := x
	for dsu.parent[root] != root {
		root = dsu.parent[root]
	}

	// path compression
	for x != root {
		p := dsu.parent[x]
		dsu.parent[x] = root
		x = p
	}

	return root
}

// Union merges the sets containing x and y.
// Returns true if the sets were separate and are now merged.
func (dsu *DSU[T]) Union(x, y T) bool {
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
func (dsu *DSU[T]) Connected(x, y T) bool {
	return dsu.Find(x) == dsu.Find(y)
}

// Groups returns a map from root -> slice of elements in that set.
func (dsu *DSU[T]) Groups() map[T][]T {
	groups := make(map[T][]T)
	for x := range dsu.parent {
		root := dsu.Find(x)
		groups[root] = append(groups[root], x)
	}
	return groups
}

// Compile-time assertion that DSU[int] implements gdsu.DSU[int].
var _ gdsu.DSU[int] = (*DSU[int])(nil)
