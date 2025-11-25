package sparse

import (
	"fmt"
	"math/rand"
	"testing"
)

// Number of elements used for benchmarking.
const NumElements = 100_000

// Simple 3D point type to benchmark generic struct performance.
type Point3D struct {
	X, Y, Z int
}

// BenchmarkSparseConstruct measures the cost of constructing a sparse DSU.
func BenchmarkSparseConstruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer() // avoid counting map allocation of keys slice
		keys := make([]int, NumElements)
		for j := 0; j < NumElements; j++ {
			keys[j] = j
		}
		b.StartTimer()

		_ = New(keys...)
	}
}

// BenchmarkSparseUnionInt benchmarks unions with int keys.
func BenchmarkSparseUnionInt(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		dsu := New[int]()
		b.StartTimer()
		for i := 0; i < NumElements; i++ {
			dsu.Union(i, i+1)
		}
	}
}

// BenchmarkSparseUnionString benchmarks unions with string keys.
func BenchmarkSparseUnionString(b *testing.B) {
	b.ReportAllocs()

	keys := make([]string, NumElements)
	for i := 0; i < NumElements; i++ {
		keys[i] = fmt.Sprintf("key-%d", i)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		dsu := New[string]()
		b.StartTimer()
		for i := 0; i < NumElements-1; i++ {
			dsu.Union(keys[i], keys[i+1])
		}
	}
}

// BenchmarkSparseUnionStruct benchmarks unions with point structs.
func BenchmarkSparseUnionStruct(b *testing.B) {
	b.ReportAllocs()

	points := make([]Point3D, NumElements)
	for i := 0; i < NumElements; i++ {
		points[i] = Point3D{i, i + 1, i + 2}
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		dsu := New[Point3D]()
		b.StartTimer()
		for i := 0; i < NumElements-1; i++ {
			dsu.Union(points[i], points[i+1])
		}
	}
}

// BenchmarkSparseFind measures Find() performance after unions.
func BenchmarkSparseFind(b *testing.B) {
	dsu := New[int]()
	for i := 0; i < NumElements; i++ {
		dsu.Union(i, i+1)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = dsu.Find(rand.Intn(NumElements))
	}
}

// BenchmarkSparseConnected tests repeated Connected() queries.
func BenchmarkSparseConnected(b *testing.B) {
	dsu := New[int]()
	for i := 0; i < NumElements; i++ {
		dsu.Union(i, i+1)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = dsu.Connected(rand.Intn(NumElements), rand.Intn(NumElements))
	}
}

// BenchmarkSparseMixedOps simulates real-world mixed operations.
func BenchmarkSparseMixedOps(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	dsu := New[int]()
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

// BenchmarkSparseGroupsSmall benchmarks Groups() with smaller size.
func BenchmarkSparseGroupsSmall(b *testing.B) {
	const small = 1_000
	dsu := New[int]()
	for i := 0; i < small; i++ {
		dsu.Union(i, i+1)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = dsu.Groups()
	}
}
