package bubble_test

import (
	"github.com/denstepanov/algostructs/src/test"
	"testing"

	"github.com/denstepanov/algostructs/src/algorithm/sort/bubble"
)

const sortType = "Bubble"

func TestBubbleSortWith10kElements(t *testing.T) {
	test.Sort(sortType, 10000, bubble.Sort, t)
}

func TestBubbleSortWithOneElement(t *testing.T) {
	test.Sort(sortType, 1, bubble.Sort, t)
}

func TestBubbleSortWothZeroElements(t *testing.T) {
	test.Sort(sortType, 0, bubble.Sort, t)
}
