package test

import (
	"github.com/denstepanov/algostructs/src/util"
	"testing"
)

// TODO: Мб получится сделать что-то более общее чем метод, привязанный к сортировкам.
func Sort(name string, sliceSize int, sort func([]int), t *testing.T) {
	slice := createIntSlice(sliceSize)
	sort(slice.Disordered)

	if !slice.AreEqual(slice.Ordered, slice.Disordered) {
		t.Fatalf("%s sort %s", name, util.MethodDoesNotWork)
	}
}
