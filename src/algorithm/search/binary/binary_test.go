package binary_test

import (
	"github.com/denstepanov/algostructs/src/algorithm/search/binary"
	"github.com/denstepanov/algostructs/src/util/slice"
	"testing"
)

// TODO: Написать больше тестов для проверки работы алгоритма.
func TestBinarySearch(t *testing.T) {
	ordered := slice.CreateOrderedUniqueSlice(10000)
	elem := 734

	result := binary.Search(ordered, ordered[elem])

	if ordered[result] != ordered[elem] {
		t.Fatal("Binary search doesn't work correctly.")
	}
}
