package heap

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

func TestHeap(t *testing.T) {
	n := 30
	vals := rand.Perm(n)
	sorted := make([]int, n)
	isorted := make([]int, n)

	for i := 0; i < n; i++ {
		sorted[i] = vals[i]
	}
	sort.Ints(sorted)
	for i := 0; i < n; i++ {
		isorted[n-1-i] = sorted[i]
	}

	minh := NewHeap(n, -1)
	maxh := NewHeap(n, 1)

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", sorted[i])
	}
	fmt.Printf("\n")
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", isorted[i])
	}
	fmt.Printf("\n")

	fmt.Println("Initilization complete. Running tests...")

	for i := 0; i < n; i++ {
		minh.Insert(vals[i])
		maxh.Insert(vals[i])
	}

	for i := 0; i < n; i++ {
		min := minh.Extract()
		max := maxh.Extract()

		fmt.Printf("minh: %d x actual: %d\n", min, sorted[i])
		fmt.Printf("maxh: %d x actual: %d\n", max, isorted[i])
		if min != sorted[i] {
			//fmt.Println("Min-Heap failed.")
			t.Fail()
		}
		if max != isorted[i] {
			//fmt.Println("Max-Heap failed.")
			t.Fail()
		}
	}
}
