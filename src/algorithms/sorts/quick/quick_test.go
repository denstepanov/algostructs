package quick_test

import (
	"github.com/denstepanov/algostructs/src/utils/slices"
	"testing"

	"github.com/denstepanov/algostructs/src/algorithms/sorts/quick"
)

func TestQuickSort(t *testing.T) {
	ordered := slices.CreateOrderedUniqueSlice(100000)
	slice := slices.ShuffleSlice(ordered)

	quick.Sort(slice, 0, len(slice)-1)

	if !slices.AreEqual(ordered, slice) {
		t.Fatal("Quick sort doesn't work correctly.")
	}
}
