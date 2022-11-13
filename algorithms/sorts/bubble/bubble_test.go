package bubble_test

import (
	"testing"

	"github.com/denstepanov/algostructs/algorithms"
	"github.com/denstepanov/algostructs/algorithms/sorts/bubble"
)

func TestDisorderedSlice(t *testing.T) {
	ordered := algorithms.GenOrderedSlice(10000)
	slice := algorithms.GenStirredSlice(ordered)

	bubble.Sort(slice)

	if !algorithms.SlicesAreEqual(ordered, slice) {
		t.Fatal("Bubble sort doesn't work correctly.")
	}
}
