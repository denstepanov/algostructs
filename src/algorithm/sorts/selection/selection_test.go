package selection_test

import (
	"github.com/denstepanov/algostructs/src/test"
	"testing"

	"github.com/denstepanov/algostructs/src/algorithm/sorts/selection"
)

const sortType = "Selection"

func TestSelectionSort(t *testing.T) {
	test.Sort(sortType, 10000, selection.Sort, t)
}
