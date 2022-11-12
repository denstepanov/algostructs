package selection_test

import (
	"testing"

	"github.com/denstepanov/algostructs/algorithms"
	"github.com/denstepanov/algostructs/algorithms/sorts/selection"
)

func TestDisorderedSlice(t *testing.T) {
	ordered := algorithms.GenOrderedSlice(200)
	slice := algorithms.GenStirredSlice(ordered)

	selection.Sort(slice)

	if !algorithms.SlicesAreEqual(ordered, slice) {
		t.Fatal("Insertion sort doesn't work correctly.")
	}
}
