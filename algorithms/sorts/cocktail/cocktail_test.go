package cocktail_test

import (
	"testing"

	"github.com/denstepanov/algostructs/algorithms"
	"github.com/denstepanov/algostructs/algorithms/sorts/cocktail"
)

func TestDisorderedSlice(t *testing.T) {
	ordered := algorithms.GenOrderedSlice(10000)
	slice := algorithms.GenStirredSlice(ordered)

	cocktail.Sort(slice)

	if !algorithms.SlicesAreEqual(ordered, slice) {
		t.Fatal("Cocktail sort doesn't work correctly.")
	}
}
