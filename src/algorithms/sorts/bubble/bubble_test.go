package bubble_test

import (
	"github.com/denstepanov/algostructs/src/test"
	"testing"

	"github.com/denstepanov/algostructs/src/algorithms/sorts/bubble"
)

func TestBubbleSort(t *testing.T) {
	test.TestSort("Bubble", bubble.Sort, t)
}
