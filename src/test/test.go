package test

import (
	"github.com/denstepanov/algostructs/src/utils"
	"testing"
)

func TestSort(name string, sort func([]int), t *testing.T) {
	slice := createIntSlice()
	sort(slice.Disordered)

	if !utils.SlicesAreEqual(slice.Ordered, slice.Disordered) {
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
