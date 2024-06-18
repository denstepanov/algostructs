package test

import (
	"github.com/denstepanov/algostructs/src/utils"
	"github.com/denstepanov/algostructs/src/utils/slices"
	"testing"
)

// TODO: Мб получится сделать что-то более общее чем метод, привязанный к сортировкам.
func Sort(name string, sliceSize int, sort func([]int), t *testing.T) {
	slice := createIntSlice(sliceSize)
	sort(slice.Disordered)

	if !slices.AreEqual(slice.Ordered, slice.Disordered) {
		t.Fatalf("%s sort %s", name, utils.MethodDoesNotWork)
	}
}
