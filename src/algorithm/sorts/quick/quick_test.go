package quick_test

import (
	"github.com/denstepanov/algostructs/src/util/slices"
	"testing"

	"github.com/denstepanov/algostructs/src/algorithm/sorts/quick"
)

func TestQuickSort(t *testing.T) {
	ord := slices.CreateOrderedUniqueSlice(100000)
	res := slices.ShuffleSlice(ord)

	quick.Sort(res)

	if slices.AreEqual(ord, res) {
		t.Fatal("Quick sorts doesn't work correctly.")
	}
}
