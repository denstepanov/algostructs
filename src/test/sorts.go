package test

import (
	"github.com/denstepanov/algostructs/src/util"
	"github.com/denstepanov/algostructs/src/util/slices"
	"testing"
)

// TODO: Мб получится сделать что-то более общее чем метод, привязанный к сортировкам.
func Sort(name string, sliceSize int, sort func([]int), t *testing.T) {
	s := createIntSlice(sliceSize)
	sort(s.Disordered)

	if !slices.AreEqual(s.Ordered, s.Disordered) {
		t.Fatalf("%s sorts %s", name, util.MethodDoesNotWork)
	}
}
