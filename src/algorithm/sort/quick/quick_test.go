package quick_test

import (
	"github.com/denstepanov/algostructs/src/util/slice"
	"testing"

	"github.com/denstepanov/algostructs/src/algorithm/sort/quick"
)

func TestQuickSort(t *testing.T) {
	ordered := slice.CreateOrderedUniqueSlice(100000)
	slice := slice.ShuffleSlice(ordered)

	quick.Sort(slice)

	if !slice.AreEqual(ordered, slice) {
		t.Fatal("Quick sort doesn't work correctly.")
	}
}
