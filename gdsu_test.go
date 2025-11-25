package gdsu

import "testing"

// fakeDSU is a trivial in-memory implementation used only to verify that
// the DSU interface is implementable as expected.
type fakeDSU[T comparable] struct{}

func (f *fakeDSU[T]) Find(x T) T                { return x }
func (f *fakeDSU[T]) Union(x, y T) bool         { return true }
func (f *fakeDSU[T]) Connected(x, y T) bool     { return true }
func (f *fakeDSU[T]) Groups() map[T][]T         { return map[T][]T{} }

func TestInterfaceCompiles(t *testing.T) {
	// This is a compile-time assertion that fakeDSU[int] satisfies DSU[int].
	var _ DSU[int] = &fakeDSU[int]{}
}

