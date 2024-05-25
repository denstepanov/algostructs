package merge_test

import (
	"github.com/denstepanov/algostructs/src/test"
	"testing"

	"github.com/denstepanov/algostructs/src/algorithms/sorts/merge"
)

func TestMergeSort(t *testing.T) {
	test.TestSort("Merge", merge.Sort, t)
}
