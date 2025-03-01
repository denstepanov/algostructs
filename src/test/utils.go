package test

import "github.com/denstepanov/algostructs/src/util/slices"

func createIntSlice(size int) SliceStruct {
	ordered := slices.CreateOrderedUniqueSlice(size)
	return SliceStruct{
		Ordered:    ordered,
		Disordered: slices.ShuffleSlice(ordered),
	}
}

type SliceStruct struct {
	Ordered, Disordered []int
}
