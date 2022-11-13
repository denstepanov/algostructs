package quick_test

import (
	"testing"

	"github.com/denstepanov/algostructs/algorithms"
	"github.com/denstepanov/algostructs/algorithms/sorts/quick"
)

func TestDisorderedSlice(t *testing.T) {
	ordered := algorithms.GenOrderedSlice(100000)
	slice := algorithms.GenStirredSlice(ordered)

	quick.Sort(slice, 0, len(slice)-1)

	if !algorithms.SlicesAreEqual(ordered, slice) {
		t.Fatal("Quick sort doesn't work correctly.")
	}
}
