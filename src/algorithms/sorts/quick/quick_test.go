package quick_test

import (
	"github.com/denstepanov/algostructs/src/utils"
	"github.com/denstepanov/algostructs/src/utils/comparsion"
	"testing"

	"github.com/denstepanov/algostructs/src/algorithms/sorts/quick"
)

func TestQuickSort(t *testing.T) {
	ordered := utils.GenOrderedSlice(100000)
	slice := utils.GenStirredSlice(ordered)

	quick.Sort(slice, 0, len(slice)-1)

	if !comparsion.SlicesAreEqual(ordered, slice) {
		t.Fatal("Quick sort doesn't work correctly.")
	}
}
