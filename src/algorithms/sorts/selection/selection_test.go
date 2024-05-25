package selection_test

import (
	"github.com/denstepanov/algostructs/src/test"
	"testing"

	"github.com/denstepanov/algostructs/src/algorithms/sorts/selection"
)

func TestSelectionSort(t *testing.T) {
	test.TestSort("Selection", selection.Sort, t)
}
