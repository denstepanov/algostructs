package test

import "github.com/denstepanov/algostructs/src/util/slice"

func createIntSlice(size int) SliceData {
	ordered := slice.CreateOrderedUniqueSlice(size)
	return SliceData{
		Ordered:    ordered,
		Disordered: slice.ShuffleSlice(ordered),
	}
}

type SliceData struct {
	Ordered, Disordered []int
}
