package search_test

import (
	"testing"

	"github.com/denstepanov/algostructs/algorithms"
	"github.com/denstepanov/algostructs/algorithms/search"
)

func TestBinarySearch(t *testing.T) {
	ordered := algorithms.GenOrderedSlice(10000)
	elem := 734

	result := search.Search(ordered, 0, len(ordered)-1, ordered[elem])

	if ordered[result] != ordered[elem] {
		t.Fatal("Binary search doesn't work correctly.")
	}
}
