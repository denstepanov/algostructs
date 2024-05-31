package quick_test

import (
	"github.com/denstepanov/algostructs/src/utils"
	"testing"

	"github.com/denstepanov/algostructs/src/algorithms/sorts/quick"
)

func TestQuickSort(t *testing.T) {
	ordered := utils.GenOrderedSlice(100000)
	slice := utils.GenStirredSlice(ordered)

	quick.Sort(slice, 0, len(slice)-1)

	if !utils.SlicesAreEqual(ordered, slice) {
		t.Fatal("Quick sort doesn't work correctly.")
	}
}
