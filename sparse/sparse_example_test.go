package sparse

import (
	"fmt"
	"sort"
)

// Example demonstrates basic usage of the sparse DSU with strings.
func Example_string() {
	dsu := New[string]()

	dsu.Union("apple", "banana")
	dsu.Union("banana", "cherry")

	// Output:
}

// Example_point shows constructing a DSU with user-defined struct keys.
func Example_point() {
	type Point struct {
		X int
		Y int
	}

	dsu := New[Point]()

	a := Point{1, 1}
	b := Point{1, 2}

	dsu.Union(a, b)

	// Output:
}

// ExampleNewpoint_empty shows constructing an empty sparse DSU
// of user-defined struct keys.
func ExampleNew_pointempty() {
	type Point struct {
		X int
		Y int
	}

	_ = New[Point]()

	// Output:
}

// ExampleNew_pointinitialized shows constructing a sparse DSU
// of user-defined struct keys initialized with a few objects.
func ExampleNew_pointinitialized() {
	type Point struct {
		X int
		Y int
	}

	a := Point{1, 1}
	b := Point{1, 2}
	c := Point{2, 3}

	_ = New(a, b, c)

	// Output:
}

// ExampleDSU_Find illustrates that Find grows the structure dynamically,
// and returns the root after unions.
func ExampleDSU_Find() {
	dsu := New[int]()

	var result int

	result = dsu.Find(33) // new element auto-created
	fmt.Println(result)

	result = dsu.Find(99)
	fmt.Println(result)

	_ = dsu.Union(33, 99)

	result = dsu.Find(99)
	fmt.Println(result)

	// Output:
	// 33
	// 99
	// 33
}

// ExampleDSU_Union illustrates connecting (union) components,
// and the return value capturing whether merge was actually necessary or not.
func ExampleDSU_Union() {
	dsu := New[int]()

	var result int
	var merged bool

	result = dsu.Find(10) // new element auto-created
	fmt.Println(result)   // expected 10

	result = dsu.Find(20)
	fmt.Println(result) // expected 20

	merged = dsu.Union(10, 20)
	fmt.Println(merged) // merge necessary, expected true

	merged = dsu.Union(20, 35)
	fmt.Println(merged) // merge necessary, expected true

	merged = dsu.Union(10, 35)
	fmt.Println(merged) // merge not necessary (due to above), expected false

	// Output:
	// 10
	// 20
	// true
	// true
	// false
}

// ExampleDSU_Connected illustrates checking connectedness between members.
func ExampleDSU_Connected() {
	dsu := New[int]()

	var result int
	var connected bool

	result = dsu.Find(10) // new element (10) auto-created
	fmt.Println(result)   // expected 10

	connected = dsu.Connected(10, 20) // element 20 auto created
	fmt.Println(connected)            // expected false

	result = dsu.Find(20)
	fmt.Println(result) // expected 20

	_ = dsu.Union(20, 10)

	connected = dsu.Connected(10, 20)
	fmt.Println(connected) // expected true

	// Output:
	// 10
	// false
	// 20
	// true
}

// ExampleDSU_Groups illustrates retrieving the connected elements as groups.
func ExampleDSU_Groups() {
	dsu := New("mozart", "bach", "gauss", "euler")

	dsu.Union("mozart", "bach")
	dsu.Union("beethoven", "bach")
	dsu.Union("mozart", "barman")

	dsu.Union("fermat", "ramanujan")
	dsu.Union("gauss", "euler")
	dsu.Union("gauss", "fermat")

	dsu.Union("gallileo", "newton")
	dsu.Union("newton", "einstein")
	dsu.Union("einstein", "bose")

	groups := dsu.Groups()
	for _, elements := range groups {
		// Elements in a line needs a predictable order for testing
		sort.Strings(elements)
		for _, element := range elements {
			fmt.Printf(" %s", element)
		}
		fmt.Printf("\n")
	}

	// Unordered output:
	//  bach barman beethoven mozart
	//  euler fermat gauss ramanujan
	//  bose einstein gallileo newton
}
