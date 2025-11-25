package compact

import (
	"testing"

	"github.com/arunksaha/gdsu"
)

func TestCompactBasic(t *testing.T) {
	dsu := New(5) // elements: 0,1,2,3,4

	if dsu.Connected(0, 1) {
		t.Fatalf("expected 0 and 1 to be initially disconnected")
	}

	dsu.Union(0, 1)
	dsu.Union(2, 3)

	if !dsu.Connected(0, 1) {
		t.Fatalf("expected 0 and 1 to be connected")
	}
	if dsu.Connected(0, 2) {
		t.Fatalf("expected 0 and 2 to be disconnected")
	}

	dsu.Union(1, 2)
	if !dsu.Connected(0, 3) {
		t.Fatalf("expected 0 and 3 to be connected after merging 1 and 2")
	}
}

func TestCompactGroups(t *testing.T) {
	dsu := New(4)
	dsu.Union(0, 1)
	dsu.Union(2, 3)

	groups := dsu.Groups()
	if len(groups) != 2 {
		t.Fatalf("expected 2 groups, got %d", len(groups))
	}
}

func TestCompactImplementsInterface(t *testing.T) {
	// Compile-time interface conformance check.
	var _ gdsu.DSU[int] = (*DSU)(nil)
}
