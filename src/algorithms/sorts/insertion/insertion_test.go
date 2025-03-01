package insertion_test

import (
	"github.com/denstepanov/algostructs/src/algorithms/sorts/insertion"
	"github.com/denstepanov/algostructs/src/test"
	"testing"
)

const sortType = "Insertion"

func TestInsertionSortWith10kElements(t *testing.T) {
	test.Sort(sortType, 10000, insertion.Sort, t)
}

func TestInsertionSortWithOneElement(t *testing.T) {
	test.Sort(sortType, 1, insertion.Sort, t)
}

func TestInsertionSortWithZeroElements(t *testing.T) {
	test.Sort(sortType, 0, insertion.Sort, t)
}
