package binary_test

import (
	"github.com/denstepanov/algostructs/src/algorithms/searches/binary"
	"github.com/denstepanov/algostructs/src/util/slices"
	"testing"
)

// TODO: Написать больше тестов для проверки работы алгоритма.
func TestBinarySearch(t *testing.T) {
	ordered := slices.CreateOrderedUniqueSlice(10000)
	elem := 734

	result := binary.Search(ordered, ordered[elem])

	if ordered[result] != ordered[elem] {
		t.Fatal("Binary searches doesn't work correctly.")
	}
}
