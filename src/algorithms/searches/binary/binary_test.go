package binary_test

import (
	"github.com/denstepanov/algostructs/src/algorithms/searches/binary"
	"github.com/denstepanov/algostructs/src/utils"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	ordered := utils.GenerateOrderedSlice(10000)
	elem := 734

	result := binary.Search(ordered, 0, len(ordered)-1, ordered[elem])

	if ordered[result] != ordered[elem] {
		t.Fatal("Binary searches doesn't work correctly.")
	}
}
