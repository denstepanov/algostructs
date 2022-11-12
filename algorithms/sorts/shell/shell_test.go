package shell_test

import (
	"testing"

	"github.com/denstepanov/algostructs/algorithms"
	"github.com/denstepanov/algostructs/algorithms/sorts/shell"
)

func TestDisorderedSlice(t *testing.T) {
	ordered := algorithms.GenOrderedSlice(200)
	slice := algorithms.GenStirredSlice(ordered)

	shell.Sort(slice)

	if !algorithms.SlicesAreEqual(ordered, slice) {
		t.Fatal("Shell sort doesn't work correctly.")
	}
}
