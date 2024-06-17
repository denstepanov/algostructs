package test

import (
	"github.com/denstepanov/algostructs/src/utils/slices"
	"testing"
)

func Sort(name string, sliceSize int, sort func([]int), t *testing.T) {
	slice := createIntSlice(sliceSize)
	sort(slice.Disordered)

	if !slices.AreEqual(slice.Ordered, slice.Disordered) {
		t.Fatalf("%s sort doesn't work correctly.", name)
	}
}

func createIntSlice(size int) SliceData {
	ordered := slices.CreateOrderedUniqueSlice(size)
	return SliceData{
		Ordered:    ordered,
		Disordered: slices.ShuffleSlice(ordered),
	}
}

type SliceData struct {
	Ordered, Disordered []int
}
