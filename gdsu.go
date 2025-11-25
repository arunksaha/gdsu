// Package gdsu defines the generic Disjoint-Set Union (Union-Find) interface.
//
// Two concrete implementations are provided in subpackages:
//
//   - sparse  – generic, map-based, no fixed capacity required.
//   - compact – int-based, slice-backed, fixed range [0, n).
//
package gdsu

// DSU is a generic Disjoint-Set Union interface.
//
// Implementations must provide:
//   - Find:      return the canonical representative (root) of x's set.
//   - Union:     merge sets containing x and y, return true if they were separate.
//   - Connected: report whether x and y belong to the same set.
//   - Groups:    return all current sets as root -> slice of elements.
type DSU[T comparable] interface {
	Find(x T) T
	Union(x, y T) bool
	Connected(x, y T) bool
	Groups() map[T][]T
}

