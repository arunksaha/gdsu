package comparison

import (
	"math/rand"
	"testing"

	"github.com/arunksaha/gdsu/compact"
	"github.com/arunksaha/gdsu/sparse"
)

const NumElements = 100_000

// BenchmarkCompareUnion compares sparse vs compact union performance.
func BenchmarkCompareUnion(b *testing.B) {
	b.Run("Sparse", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			b.StopTimer()
			dsu := sparse.New[int]()
			b.StartTimer()
			for i := 0; i < NumElements-1; i++ {
				dsu.Union(i, i+1)
			}
		}
	})

	b.Run("Compact", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			b.StopTimer()
			dsu := compact.New(NumElements)
			b.StartTimer()
			for i := 0; i < NumElements-1; i++ {
				dsu.Union(i, i+1)
			}
		}
	})
}

// BenchmarkCompareFind compares Find() between sparse and compact.
func BenchmarkCompareFind(b *testing.B) {
	b.Run("Sparse", func(b *testing.B) {
		dsu := sparse.New[int]()
		for i := 0; i < NumElements-1; i++ {
			dsu.Union(i, i+1)
		}
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = dsu.Find(rand.Intn(NumElements))
		}
	})

	b.Run("Compact", func(b *testing.B) {
		dsu := compact.New(NumElements)
		for i := 0; i < NumElements-1; i++ {
			dsu.Union(i, i+1)
		}
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = dsu.Find(rand.Intn(NumElements))
		}
	})
}
