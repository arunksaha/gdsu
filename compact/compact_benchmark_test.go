package compact

import (
	"math/rand"
	"testing"
)

const NumElements = 100_000

// BenchmarkCompactConstruct measures the cost of constructing a compact DSU.
func BenchmarkCompactConstruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// There is no external key generation needed; just exclude the loop/iteration overhead.
		b.StartTimer()
		_ = New(NumElements)
		b.StopTimer()
	}
}

// BenchmarkCompactUnion benchmarks sequential unions.
func BenchmarkCompactUnion(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		dsu := New(NumElements)
		b.StartTimer()
		for i := 0; i < NumElements-1; i++ {
			dsu.Union(i, i+1)
		}
	}
}

// BenchmarkCompactFind tests repeated Find() operations.
func BenchmarkCompactFind(b *testing.B) {
	dsu := New(NumElements)
	for i := 0; i < NumElements-1; i++ {
		dsu.Union(i, i+1)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = dsu.Find(rand.Intn(NumElements))
	}
}

// BenchmarkCompactConnected tests repeated Connected() queries.
func BenchmarkCompactConnected(b *testing.B) {
	dsu := New(NumElements)
	for i := 0; i < NumElements-1; i++ {
		dsu.Union(i, i+1)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = dsu.Connected(rand.Intn(NumElements), rand.Intn(NumElements))
	}
}

// BenchmarkCompactMixedOps runs random unions and queries.
func BenchmarkCompactMixedOps(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	dsu := New(NumElements)
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		x := rand.Intn(NumElements)
		y := rand.Intn(NumElements)
		if rand.Intn(2) == 0 {
			dsu.Union(x, y)
		} else {
			_ = dsu.Connected(x, y)
		}
	}
}

// BenchmarkCompactWorstCase builds long chains (bad initial shape).
func BenchmarkCompactWorstCase(b *testing.B) {
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		b.StopTimer()
		dsu := New(NumElements)
		b.StartTimer()
		for i := 0; i < NumElements-1; i++ {
			dsu.Union(i, i+1) // creates a degenerate chain before compression
		}
	}
}
