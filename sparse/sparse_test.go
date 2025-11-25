package sparse

import (
	"testing"

	"github.com/arunksaha/gdsu"
)

func TestSparseBasicInt(t *testing.T) {
	dsu := New([]int{1, 2, 3, 4, 5})

	if dsu.Connected(1, 2) {
		t.Fatalf("expected 1 and 2 to be initially disconnected")
	}

	dsu.Union(1, 2)
	dsu.Union(3, 4)

	if !dsu.Connected(1, 2) {
		t.Fatalf("expected 1 and 2 to be connected")
	}
	if dsu.Connected(1, 3) {
		t.Fatalf("expected 1 and 3 to be disconnected")
	}

	dsu.Union(2, 3)
	if !dsu.Connected(1, 4) {
		t.Fatalf("expected 1 and 4 to be connected after merging 2 and 3")
	}
}

func TestSparseGenericString(t *testing.T) {
	dsu := New([]string{"a", "b", "c", "d"})

	dsu.Union("a", "b")
	dsu.Union("c", "d")

	if !dsu.Connected("a", "b") {
		t.Fatalf("expected a and b to be connected")
	}
	if dsu.Connected("a", "c") {
		t.Fatalf("expected a and c to be disconnected")
	}

	groups := dsu.Groups()
	if len(groups) != 2 {
		t.Fatalf("expected 2 groups, got %d", len(groups))
	}
}

func TestSparseImplementsInterface(t *testing.T) {
	// Compile-time interface conformance check.
	var _ gdsu.DSU[int] = (*DSU[int])(nil)
	var _ gdsu.DSU[string] = (*DSU[string])(nil)
}
