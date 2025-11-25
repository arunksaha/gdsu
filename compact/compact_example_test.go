package compact

import (
	"fmt"
	"sort"
)

// Example demonstrates basic usage of the compact DSU with ints.
func Example() {
	dsu := New(5) // elements 0..4

	dsu.Union(0, 1)
	dsu.Union(3, 4)

	fmt.Println(dsu.Connected(0, 1))
	fmt.Println(dsu.Connected(0, 3))
	fmt.Println(len(dsu.Groups()))

	// Output:
	// true
	// false
	// 3
}

// ExampleNew_range_empty shows constructing an empty compact DSU (size 0).
func ExampleNew_range_empty() {
	dsu := New(0)

	fmt.Println(len(dsu.Groups()))

	// Output:
	// 0
}

// ExampleNew_range_initialized shows constructing a compact DSU
// pre-sized for a range of ints.
func ExampleNew_range_initialized() {
	dsu := New(4) // elements 0..3

	fmt.Println(len(dsu.Groups()))

	// Output:
	// 4
}

// ExampleDSU_Find illustrates Find after a few unions.
func ExampleDSU_Find() {
	dsu := New(3)

	fmt.Println(dsu.Find(1)) // should be itself, i.e., 1

	dsu.Union(0, 1)

	fmt.Println(dsu.Find(1)) // should compress to root 0

	// Output:
	// 1
	// 0
}

// ExampleDSU_Union illustrates connecting components and the returned merge flag.
func ExampleDSU_Union() {
	dsu := New(6)

	fmt.Println(dsu.Union(1, 2)) // merge needed
	fmt.Println(dsu.Union(2, 5)) // merge needed
	fmt.Println(dsu.Union(1, 5)) // already connected
	fmt.Println(dsu.Union(3, 4)) // merge needed

	// Output:
	// true
	// true
	// false
	// true
}

// ExampleDSU_Connected illustrates checking connectedness between members.
func ExampleDSU_Connected() {
	dsu := New(4)

	fmt.Println(dsu.Connected(0, 1)) // initially disconnected

	dsu.Union(0, 1)

	fmt.Println(dsu.Connected(0, 1)) // now connected
	fmt.Println(dsu.Connected(1, 2)) // still disconnected

	// Output:
	// false
	// true
	// false
}

// ExampleDSU_Groups illustrates retrieving the connected elements as groups.
func ExampleDSU_Groups() {
	dsu := New(100) // elements 0..99

	dsu.Union(3, 9)
	dsu.Union(27, 81)
	dsu.Union(3, 27)

	dsu.Union(2, 64)
	dsu.Union(4, 16)
	dsu.Union(32, 8)
	dsu.Union(4, 32)
	dsu.Union(16, 2)

	dsu.Union(5, 25)

	groups := dsu.Groups()

	for _, elements := range groups {
		// Skip single element sets
		if len(elements) == 1 {
			continue
		}
		// Elements in a line needs a predictable order for testing
		sort.Ints(elements)
		separator := ""
		for _, element := range elements {
			fmt.Printf("%s%d", separator, element)
			separator = ","
		}
		fmt.Printf("\n")
	}

	// Unordered output:
	// 2,4,8,16,32,64
	// 3,9,27,81
	// 5,25
}
