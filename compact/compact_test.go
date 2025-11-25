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

// TestNewNegativeSize ensures New() dos not panic when called with a negative size.
func TestCompactNewNegativeSize(t *testing.T) {
	_ = New(-5)
}

// TestFindOutOfBounds ensures Find() panics when index is outside the valid range.
func TestCompactFindOutOfBounds(t *testing.T) {
	dsu := New(5)

	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic on out-of-range Find(), got none")
		}
	}()
	_ = dsu.Find(10) // invalid index
}

// TestCompactUnionOutOfBounds ensures Find() panics when index is outside the valid range.
func TestCompactUnionOutOfBounds(t *testing.T) {
	dsu := New(5)

	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic on out-of-range Find(), got none")
		}
	}()
	_ = dsu.Union(3, 10) // invalid index
}

// TestCompactConnectedOutOfBounds ensures Find() panics when index is outside the valid range.
func TestCompactConnectedOutOfBounds(t *testing.T) {
	dsu := New(5)

	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic on out-of-range Find(), got none")
		}
	}()
	_ = dsu.Connected(10, 3) // invalid index
}

// TestUnionAlreadyConnected ensures Union() returns false when x and y are already in the same set.
func TestCompactUnionAlreadyConnected(t *testing.T) {
	dsu := New(5)

	dsu.Union(1, 2)
	if !dsu.Connected(1, 2) {
		t.Fatal("setup error: 1 and 2 should be connected")
	}

	// Calling Union again should indicate "no merge happened"
	if dsu.Union(1, 2) {
		t.Fatal("Union should return false when merging an already-connected pair")
	}
}

// TestUnionRankSwap ensures union-by-rank attaches the lower-rank root under the higher-rank root.
func TestCompactUnionRankSwap(t *testing.T) {
	dsu := New(5)

	// Step 1: merge 1 and 2 → root(2) becomes 1 (rank increases)
	dsu.Union(1, 2)

	// Step 2: artificially increase rank of root(1) by merging 1 and 3
	dsu.Union(1, 3)

	// At this point, root(1) has higher rank than root(4) (rank 0)
	rootBefore := dsu.Find(1)

	// Step 3: merge 4 into the existing high-rank set
	dsu.Union(4, 1)

	rootAfter := dsu.Find(4)

	if rootAfter != rootBefore {
		t.Fatalf("expected node 4 to attach under higher-rank root %d, got %d", rootBefore, rootAfter)
	}
}

func TestCompactImplementsInterface(t *testing.T) {
	// Compile-time interface conformance check.
	var _ gdsu.DSU[int] = (*DSU)(nil)
}
