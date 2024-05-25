package test

import (
	"github.com/denstepanov/algostructs/src/utils"
	"github.com/denstepanov/algostructs/src/utils/comparsion"
	"testing"
)

func TestSort(name string, sort func([]int), t *testing.T) {
	slice := createIntSlice()
	sort(slice.Disordered)

	if !comparsion.SlicesAreEqual(slice.Ordered, slice.Disordered) {
		t.Fatalf("%s sort doesn't work correctly.", name)
	}
}

func createIntSlice() SliceData {
	ordered := utils.GenOrderedSlice(10000)
	return SliceData{
		Ordered:    ordered,
		Disordered: utils.GenStirredSlice(ordered),
	}
}

type SliceData struct {
	Ordered, Disordered []int
}
