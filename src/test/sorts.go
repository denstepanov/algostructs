package test

import (
	"github.com/denstepanov/algostructs/src/utils"
	"testing"
)

func Sort(name string, sliceSize int, sort func([]int), t *testing.T) {
	slice := createIntSlice(sliceSize)
	sort(slice.Disordered)

	if !utils.SlicesAreEqual(slice.Ordered, slice.Disordered) {
		t.Fatalf("%s sort doesn't work correctly.", name)
	}
}

func createIntSlice(size int) SliceData {
	ordered := utils.GenerateOrderedSlice(size)
	return SliceData{
		Ordered:    ordered,
		Disordered: utils.ShuffleSlice(ordered),
	}
}

type SliceData struct {
	Ordered, Disordered []int
}
