package merge_test

import (
	"github.com/denstepanov/algostructs/src/test"
	"testing"

	"github.com/denstepanov/algostructs/src/algorithm/sort/merge"
)

const sortType = "Merge"

func TestMergeSortWith10kElements(t *testing.T) {
	test.Sort(sortType, 10000, merge.Sort, t)
}

func TestMergeSortWithOneElement(t *testing.T) {
	test.Sort(sortType, 10000, merge.Sort, t)
}

func TestMergeSortWithZeroElements(t *testing.T) {
	test.Sort(sortType, 10000, merge.Sort, t)
}
