package insertion_test

import (
	"github.com/denstepanov/algostructs/src/algorithms/sorts/insertion"
	"github.com/denstepanov/algostructs/src/test"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	test.Sort("Insertion", insertion.Sort, t)
}
