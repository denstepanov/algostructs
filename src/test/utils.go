package test

import "github.com/denstepanov/algostructs/src/utils/slices"

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
