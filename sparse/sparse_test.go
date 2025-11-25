package sparse

import (
	"testing"

	"github.com/arunksaha/gdsu"
)

// TestSparseBasicInt tests basic functionality using int elements.
func TestSparseBasicInt(t *testing.T) {
	dsu := New(1, 2, 3, 4, 5)

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

// TestSparseBasicInt tests basic functionality using string elements.
func TestSparseGenericString(t *testing.T) {
	dsu := New("a", "b", "c", "d")

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

// TestSparseFindCreatesNew checks that Find() auto-creates unseen elements.
func TestSparseFindCreatesNew(t *testing.T) {
	dsu := New[string]()
	root := dsu.Find("x") // should not panic, should auto-add
	if root != "x" {
		t.Fatalf("expected root to be x, got %s", root)
	}
}

// TestSparseUnionOnNewElements ensures Union() works even if elements were never added before.
func TestSparseUnionOnNewElements(t *testing.T) {
	dsu := New[int]()
	merged := dsu.Union(10, 20)
	if !merged {
		t.Fatalf("expected merge to return true")
	}
	if !dsu.Connected(10, 20) {
		t.Fatalf("10 and 20 should be connected")
	}
}

// TestSparseIdempotentUnion ensures repeated unions do not break structure.
func TestSparseIdempotentUnion(t *testing.T) {
	dsu := New[int]()
	dsu.Union(1, 2)
	dsu.Union(1, 2)
	dsu.Union(2, 1)

	if !dsu.Connected(1, 2) {
		t.Fatalf("1 and 2 should remain connected")
	}
	if len(dsu.Groups()) != 1 {
		t.Fatalf("expected 1 group, got %d", len(dsu.Groups()))
	}
}

// TestSparseMultipleGroups checks if multiple disjoint groups are maintained correctly.
func TestSparseMultipleGroups(t *testing.T) {
	dsu := New[string]()
	dsu.Union("a", "b")
	dsu.Union("c", "d")

	if dsu.Connected("a", "c") {
		t.Fatalf("a and c should not be connected")
	}

	groups := dsu.Groups()
	if len(groups) != 2 {
		t.Fatalf("expected 2 groups, got %d", len(groups))
	}
}

// TestSparseDeepChainPathCompression validates that long chains compress correctly.
func TestSparseDeepChainPathCompression(t *testing.T) {
	dsu := New[int]()
	for i := 1; i <= 10; i++ {
		dsu.Union(i, i+1)
	}
	rootBefore := dsu.Find(10)
	rootAfter := dsu.Find(1)

	if rootBefore != rootAfter {
		t.Fatalf("path compression failed, roots differ")
	}
}

// TestSparseUnionRankSwapSparse verifies that union-by-rank correctly attaches
// the lower-rank tree under the higher-rank tree in the sparse DSU.
func TestSparseUnionRankSwapSparse(t *testing.T) {
	d := New("a", "b", "c")

	// Create: a <- b (rank[a] = 1)
	d.Union("a", "b")

	// Without interfering with a or b, union c under b so that:
	// Find(b) == 'a' and rank[a] > rank[c]
	rootC := d.Find("c")
	if rootC != "c" {
		t.Fatalf("expected c to be its own root, got %v", rootC)
	}

	// Now union c with a (rank[a] = 1, rank[c] = 0),
	// so c should attach under a, NOT the other way around.
	d.Union("c", "a")

	// Check representatives
	if !d.Connected("b", "c") {
		t.Fatalf("expected b and c to be connected after union-by-rank")
	}

	root := d.Find("c")
	if root != "a" {
		t.Fatalf("expected a to remain the root due to higher rank, got %v", root)
	}
}

// TestSparseImplementsInterface tests interface compliance.
func TestSparseImplementsInterface(t *testing.T) {
	// Compile-time interface conformance check.
	var _ gdsu.DSU[int] = (*DSU[int])(nil)
	var _ gdsu.DSU[string] = (*DSU[string])(nil)
}
